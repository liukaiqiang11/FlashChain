package core

import (
	pb "fchain/proto"
)

func CreateEnvelopeGroup(group *TransactionGroup, envs []*pb.Envelope) []*pb.Envelopes {
	var envGroup []*pb.Envelopes
	var envelopes = &pb.Envelopes{}
	var env *pb.Envelope
	var num = 10
	for _, g := range group.Group {
		if len(g) != 0 {
			cg := CreateConflictGraph(g)
			cg.TopologicalSort()
			for _, str := range cg.Queue {
				env, envs = GetEnvelope(str, envs)
				if env != nil {
					envelopes.Envelope = append(envelopes.Envelope, env)
				}
			}
			envGroup = append(envGroup, envelopes)
			envelopes = &pb.Envelopes{}
		}
	}
	for len(envs) != 0 {
		if len(envs) <= num {
			envGroup = append(envGroup, &pb.Envelopes{Envelope: envs})
			envs = nil
		} else {
			envGroup = append(envGroup, &pb.Envelopes{Envelope: envs[:num]})
			envs = envs[num:]
		}
	}

	return envGroup
}
