// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ShoulderStabilityExerciseName uint16

const (
	ShoulderStabilityExerciseName90DegreeCableExternalRotation          ShoulderStabilityExerciseName = 0
	ShoulderStabilityExerciseNameBandExternalRotation                   ShoulderStabilityExerciseName = 1
	ShoulderStabilityExerciseNameBandInternalRotation                   ShoulderStabilityExerciseName = 2
	ShoulderStabilityExerciseNameBentArmLateralRaiseAndExternalRotation ShoulderStabilityExerciseName = 3
	ShoulderStabilityExerciseNameCableExternalRotation                  ShoulderStabilityExerciseName = 4
	ShoulderStabilityExerciseNameDumbbellFacePullWithExternalRotation   ShoulderStabilityExerciseName = 5
	ShoulderStabilityExerciseNameFloorIRaise                            ShoulderStabilityExerciseName = 6
	ShoulderStabilityExerciseNameWeightedFloorIRaise                    ShoulderStabilityExerciseName = 7
	ShoulderStabilityExerciseNameFloorTRaise                            ShoulderStabilityExerciseName = 8
	ShoulderStabilityExerciseNameWeightedFloorTRaise                    ShoulderStabilityExerciseName = 9
	ShoulderStabilityExerciseNameFloorYRaise                            ShoulderStabilityExerciseName = 10
	ShoulderStabilityExerciseNameWeightedFloorYRaise                    ShoulderStabilityExerciseName = 11
	ShoulderStabilityExerciseNameInclineIRaise                          ShoulderStabilityExerciseName = 12
	ShoulderStabilityExerciseNameWeightedInclineIRaise                  ShoulderStabilityExerciseName = 13
	ShoulderStabilityExerciseNameInclineLRaise                          ShoulderStabilityExerciseName = 14
	ShoulderStabilityExerciseNameWeightedInclineLRaise                  ShoulderStabilityExerciseName = 15
	ShoulderStabilityExerciseNameInclineTRaise                          ShoulderStabilityExerciseName = 16
	ShoulderStabilityExerciseNameWeightedInclineTRaise                  ShoulderStabilityExerciseName = 17
	ShoulderStabilityExerciseNameInclineWRaise                          ShoulderStabilityExerciseName = 18
	ShoulderStabilityExerciseNameWeightedInclineWRaise                  ShoulderStabilityExerciseName = 19
	ShoulderStabilityExerciseNameInclineYRaise                          ShoulderStabilityExerciseName = 20
	ShoulderStabilityExerciseNameWeightedInclineYRaise                  ShoulderStabilityExerciseName = 21
	ShoulderStabilityExerciseNameLyingExternalRotation                  ShoulderStabilityExerciseName = 22
	ShoulderStabilityExerciseNameSeatedDumbbellExternalRotation         ShoulderStabilityExerciseName = 23
	ShoulderStabilityExerciseNameStandingLRaise                         ShoulderStabilityExerciseName = 24
	ShoulderStabilityExerciseNameSwissBallIRaise                        ShoulderStabilityExerciseName = 25
	ShoulderStabilityExerciseNameWeightedSwissBallIRaise                ShoulderStabilityExerciseName = 26
	ShoulderStabilityExerciseNameSwissBallTRaise                        ShoulderStabilityExerciseName = 27
	ShoulderStabilityExerciseNameWeightedSwissBallTRaise                ShoulderStabilityExerciseName = 28
	ShoulderStabilityExerciseNameSwissBallWRaise                        ShoulderStabilityExerciseName = 29
	ShoulderStabilityExerciseNameWeightedSwissBallWRaise                ShoulderStabilityExerciseName = 30
	ShoulderStabilityExerciseNameSwissBallYRaise                        ShoulderStabilityExerciseName = 31
	ShoulderStabilityExerciseNameWeightedSwissBallYRaise                ShoulderStabilityExerciseName = 32
	ShoulderStabilityExerciseNameInvalid                                ShoulderStabilityExerciseName = 0xFFFF
)

func (s ShoulderStabilityExerciseName) Uint16() uint16 { return uint16(s) }

func (s ShoulderStabilityExerciseName) String() string {
	switch s {
	case ShoulderStabilityExerciseName90DegreeCableExternalRotation:
		return "90_degree_cable_external_rotation"
	case ShoulderStabilityExerciseNameBandExternalRotation:
		return "band_external_rotation"
	case ShoulderStabilityExerciseNameBandInternalRotation:
		return "band_internal_rotation"
	case ShoulderStabilityExerciseNameBentArmLateralRaiseAndExternalRotation:
		return "bent_arm_lateral_raise_and_external_rotation"
	case ShoulderStabilityExerciseNameCableExternalRotation:
		return "cable_external_rotation"
	case ShoulderStabilityExerciseNameDumbbellFacePullWithExternalRotation:
		return "dumbbell_face_pull_with_external_rotation"
	case ShoulderStabilityExerciseNameFloorIRaise:
		return "floor_i_raise"
	case ShoulderStabilityExerciseNameWeightedFloorIRaise:
		return "weighted_floor_i_raise"
	case ShoulderStabilityExerciseNameFloorTRaise:
		return "floor_t_raise"
	case ShoulderStabilityExerciseNameWeightedFloorTRaise:
		return "weighted_floor_t_raise"
	case ShoulderStabilityExerciseNameFloorYRaise:
		return "floor_y_raise"
	case ShoulderStabilityExerciseNameWeightedFloorYRaise:
		return "weighted_floor_y_raise"
	case ShoulderStabilityExerciseNameInclineIRaise:
		return "incline_i_raise"
	case ShoulderStabilityExerciseNameWeightedInclineIRaise:
		return "weighted_incline_i_raise"
	case ShoulderStabilityExerciseNameInclineLRaise:
		return "incline_l_raise"
	case ShoulderStabilityExerciseNameWeightedInclineLRaise:
		return "weighted_incline_l_raise"
	case ShoulderStabilityExerciseNameInclineTRaise:
		return "incline_t_raise"
	case ShoulderStabilityExerciseNameWeightedInclineTRaise:
		return "weighted_incline_t_raise"
	case ShoulderStabilityExerciseNameInclineWRaise:
		return "incline_w_raise"
	case ShoulderStabilityExerciseNameWeightedInclineWRaise:
		return "weighted_incline_w_raise"
	case ShoulderStabilityExerciseNameInclineYRaise:
		return "incline_y_raise"
	case ShoulderStabilityExerciseNameWeightedInclineYRaise:
		return "weighted_incline_y_raise"
	case ShoulderStabilityExerciseNameLyingExternalRotation:
		return "lying_external_rotation"
	case ShoulderStabilityExerciseNameSeatedDumbbellExternalRotation:
		return "seated_dumbbell_external_rotation"
	case ShoulderStabilityExerciseNameStandingLRaise:
		return "standing_l_raise"
	case ShoulderStabilityExerciseNameSwissBallIRaise:
		return "swiss_ball_i_raise"
	case ShoulderStabilityExerciseNameWeightedSwissBallIRaise:
		return "weighted_swiss_ball_i_raise"
	case ShoulderStabilityExerciseNameSwissBallTRaise:
		return "swiss_ball_t_raise"
	case ShoulderStabilityExerciseNameWeightedSwissBallTRaise:
		return "weighted_swiss_ball_t_raise"
	case ShoulderStabilityExerciseNameSwissBallWRaise:
		return "swiss_ball_w_raise"
	case ShoulderStabilityExerciseNameWeightedSwissBallWRaise:
		return "weighted_swiss_ball_w_raise"
	case ShoulderStabilityExerciseNameSwissBallYRaise:
		return "swiss_ball_y_raise"
	case ShoulderStabilityExerciseNameWeightedSwissBallYRaise:
		return "weighted_swiss_ball_y_raise"
	default:
		return "ShoulderStabilityExerciseNameInvalid(" + strconv.FormatUint(uint64(s), 10) + ")"
	}
}

// FromString parse string into ShoulderStabilityExerciseName constant it's represent, return ShoulderStabilityExerciseNameInvalid if not found.
func ShoulderStabilityExerciseNameFromString(s string) ShoulderStabilityExerciseName {
	switch s {
	case "90_degree_cable_external_rotation":
		return ShoulderStabilityExerciseName90DegreeCableExternalRotation
	case "band_external_rotation":
		return ShoulderStabilityExerciseNameBandExternalRotation
	case "band_internal_rotation":
		return ShoulderStabilityExerciseNameBandInternalRotation
	case "bent_arm_lateral_raise_and_external_rotation":
		return ShoulderStabilityExerciseNameBentArmLateralRaiseAndExternalRotation
	case "cable_external_rotation":
		return ShoulderStabilityExerciseNameCableExternalRotation
	case "dumbbell_face_pull_with_external_rotation":
		return ShoulderStabilityExerciseNameDumbbellFacePullWithExternalRotation
	case "floor_i_raise":
		return ShoulderStabilityExerciseNameFloorIRaise
	case "weighted_floor_i_raise":
		return ShoulderStabilityExerciseNameWeightedFloorIRaise
	case "floor_t_raise":
		return ShoulderStabilityExerciseNameFloorTRaise
	case "weighted_floor_t_raise":
		return ShoulderStabilityExerciseNameWeightedFloorTRaise
	case "floor_y_raise":
		return ShoulderStabilityExerciseNameFloorYRaise
	case "weighted_floor_y_raise":
		return ShoulderStabilityExerciseNameWeightedFloorYRaise
	case "incline_i_raise":
		return ShoulderStabilityExerciseNameInclineIRaise
	case "weighted_incline_i_raise":
		return ShoulderStabilityExerciseNameWeightedInclineIRaise
	case "incline_l_raise":
		return ShoulderStabilityExerciseNameInclineLRaise
	case "weighted_incline_l_raise":
		return ShoulderStabilityExerciseNameWeightedInclineLRaise
	case "incline_t_raise":
		return ShoulderStabilityExerciseNameInclineTRaise
	case "weighted_incline_t_raise":
		return ShoulderStabilityExerciseNameWeightedInclineTRaise
	case "incline_w_raise":
		return ShoulderStabilityExerciseNameInclineWRaise
	case "weighted_incline_w_raise":
		return ShoulderStabilityExerciseNameWeightedInclineWRaise
	case "incline_y_raise":
		return ShoulderStabilityExerciseNameInclineYRaise
	case "weighted_incline_y_raise":
		return ShoulderStabilityExerciseNameWeightedInclineYRaise
	case "lying_external_rotation":
		return ShoulderStabilityExerciseNameLyingExternalRotation
	case "seated_dumbbell_external_rotation":
		return ShoulderStabilityExerciseNameSeatedDumbbellExternalRotation
	case "standing_l_raise":
		return ShoulderStabilityExerciseNameStandingLRaise
	case "swiss_ball_i_raise":
		return ShoulderStabilityExerciseNameSwissBallIRaise
	case "weighted_swiss_ball_i_raise":
		return ShoulderStabilityExerciseNameWeightedSwissBallIRaise
	case "swiss_ball_t_raise":
		return ShoulderStabilityExerciseNameSwissBallTRaise
	case "weighted_swiss_ball_t_raise":
		return ShoulderStabilityExerciseNameWeightedSwissBallTRaise
	case "swiss_ball_w_raise":
		return ShoulderStabilityExerciseNameSwissBallWRaise
	case "weighted_swiss_ball_w_raise":
		return ShoulderStabilityExerciseNameWeightedSwissBallWRaise
	case "swiss_ball_y_raise":
		return ShoulderStabilityExerciseNameSwissBallYRaise
	case "weighted_swiss_ball_y_raise":
		return ShoulderStabilityExerciseNameWeightedSwissBallYRaise
	default:
		return ShoulderStabilityExerciseNameInvalid
	}
}

// List returns all constants.
func ListShoulderStabilityExerciseName() []ShoulderStabilityExerciseName {
	return []ShoulderStabilityExerciseName{
		ShoulderStabilityExerciseName90DegreeCableExternalRotation,
		ShoulderStabilityExerciseNameBandExternalRotation,
		ShoulderStabilityExerciseNameBandInternalRotation,
		ShoulderStabilityExerciseNameBentArmLateralRaiseAndExternalRotation,
		ShoulderStabilityExerciseNameCableExternalRotation,
		ShoulderStabilityExerciseNameDumbbellFacePullWithExternalRotation,
		ShoulderStabilityExerciseNameFloorIRaise,
		ShoulderStabilityExerciseNameWeightedFloorIRaise,
		ShoulderStabilityExerciseNameFloorTRaise,
		ShoulderStabilityExerciseNameWeightedFloorTRaise,
		ShoulderStabilityExerciseNameFloorYRaise,
		ShoulderStabilityExerciseNameWeightedFloorYRaise,
		ShoulderStabilityExerciseNameInclineIRaise,
		ShoulderStabilityExerciseNameWeightedInclineIRaise,
		ShoulderStabilityExerciseNameInclineLRaise,
		ShoulderStabilityExerciseNameWeightedInclineLRaise,
		ShoulderStabilityExerciseNameInclineTRaise,
		ShoulderStabilityExerciseNameWeightedInclineTRaise,
		ShoulderStabilityExerciseNameInclineWRaise,
		ShoulderStabilityExerciseNameWeightedInclineWRaise,
		ShoulderStabilityExerciseNameInclineYRaise,
		ShoulderStabilityExerciseNameWeightedInclineYRaise,
		ShoulderStabilityExerciseNameLyingExternalRotation,
		ShoulderStabilityExerciseNameSeatedDumbbellExternalRotation,
		ShoulderStabilityExerciseNameStandingLRaise,
		ShoulderStabilityExerciseNameSwissBallIRaise,
		ShoulderStabilityExerciseNameWeightedSwissBallIRaise,
		ShoulderStabilityExerciseNameSwissBallTRaise,
		ShoulderStabilityExerciseNameWeightedSwissBallTRaise,
		ShoulderStabilityExerciseNameSwissBallWRaise,
		ShoulderStabilityExerciseNameWeightedSwissBallWRaise,
		ShoulderStabilityExerciseNameSwissBallYRaise,
		ShoulderStabilityExerciseNameWeightedSwissBallYRaise,
	}
}