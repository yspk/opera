package election

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/Fantom-foundation/go-lachesis/src/hash"
	"github.com/Fantom-foundation/go-lachesis/src/inter/idx"
	"github.com/Fantom-foundation/go-lachesis/src/inter/pos"
	"github.com/Fantom-foundation/go-lachesis/src/logger"
)

// TODO implement&test coinRound
//const coinRound = 10 // every 10th round is a round with pseudorandom votes

type (
	// Election cached data of election algorithm.
	Election struct {
		// election params
		frameToDecide idx.Frame

		members pos.Members

		// election state
		decidedRoots map[common.Address]voteValue // decided roots at "frameToDecide"
		votes        map[voteId]voteValue

		// external world
		observe RootForklessCausesRootFn

		logger.Instance
	}

	// RootForklessCausesRootFn returns hash of root B, if root B forkless causes root A.
	// Due to a fork, there may be many roots B with the same slot,
	// but A may be forkless caused only by one of them (if no more than 1/3n are Byzantine), with a specific hash.
	RootForklessCausesRootFn func(a hash.Event, b common.Address, f idx.Frame) *hash.Event

	// Slot specifies a root slot {addr, frame}. Normal members can have only one root with this pair.
	// Due to a fork, different roots may occupy the same slot
	Slot struct {
		Frame idx.Frame
		Addr  common.Address
	}

	// RootAndSlot specifies concrete root of slot.
	RootAndSlot struct {
		Root hash.Event
		Slot Slot
	}
)

type voteId struct {
	fromRoot  hash.Event
	forMember common.Address
}
type voteValue struct {
	decided      bool
	yes          bool
	observedRoot hash.Event
}

type ElectionRes struct {
	Frame   idx.Frame
	Atropos hash.Event
}

func New(
	members pos.Members,
	frameToDecide idx.Frame,
	forklessCausesFn RootForklessCausesRootFn,
) *Election {
	el := &Election{
		observe: forklessCausesFn,

		Instance: logger.MakeInstance(),
	}

	el.Reset(members, frameToDecide)

	return el
}

// erase the current election state, prepare for new election frame
func (el *Election) Reset(members pos.Members, frameToDecide idx.Frame) {
	el.members = members
	el.frameToDecide = frameToDecide
	el.votes = make(map[voteId]voteValue)
	el.decidedRoots = make(map[common.Address]voteValue)
}

// return root slots which are not within el.decidedRoots
func (el *Election) notDecidedRoots() []common.Address {
	notDecidedRoots := make([]common.Address, 0, len(el.members))

	for member := range el.members {
		if _, ok := el.decidedRoots[member]; !ok {
			notDecidedRoots = append(notDecidedRoots, member)
		}
	}
	if len(notDecidedRoots)+len(el.decidedRoots) != len(el.members) { // sanity check
		el.Log.Crit("Mismatch of roots")
	}
	return notDecidedRoots
}

// observedRoots returns all the roots at the specified frame which do forkless cause the specified root.
func (el *Election) observedRoots(root hash.Event, frame idx.Frame) []RootAndSlot {
	observedRoots := make([]RootAndSlot, 0, len(el.members))
	for member := range el.members {
		slot := Slot{
			Frame: frame,
			Addr:  member,
		}
		observedRoot := el.observe(root, member, frame)
		if observedRoot != nil {
			observedRoots = append(observedRoots, RootAndSlot{
				Root: *observedRoot,
				Slot: slot,
			})
		}
	}
	return observedRoots
}
