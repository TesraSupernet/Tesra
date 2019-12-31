/*
 * Copyright (C) 2019 The TesraSupernet Authors
 * This file is part of The TesraSupernet library.
 *
 * The TesraSupernet is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The TesraSupernet is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The TesraSupernet.  If not, see <http://www.gnu.org/licenses/>.
 */

package actor

import (
	"github.com/TesraSupernet/tesraevent/actor"
	cactor "github.com/TesraSupernet/Tesra/consensus/actor"
)

var consensusSrvPid *actor.PID

func SetConsensusPid(actr *actor.PID) {
	consensusSrvPid = actr
}

//start consensus to consensus actor
func ConsensusSrvStart() error {
	if consensusSrvPid != nil {
		consensusSrvPid.Tell(&cactor.StartConsensus{})
	}
	return nil
}

//halt consensus to consensus actor
func ConsensusSrvHalt() error {
	if consensusSrvPid != nil {
		consensusSrvPid.Tell(&cactor.StopConsensus{})
	}
	return nil
}
