// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LungeExerciseName uint16

const (
	LungeExerciseNameOverheadLunge                                 LungeExerciseName = 0
	LungeExerciseNameLungeMatrix                                   LungeExerciseName = 1
	LungeExerciseNameWeightedLungeMatrix                           LungeExerciseName = 2
	LungeExerciseNameAlternatingBarbellForwardLunge                LungeExerciseName = 3
	LungeExerciseNameAlternatingDumbbellLungeWithReach             LungeExerciseName = 4
	LungeExerciseNameBackFootElevatedDumbbellSplitSquat            LungeExerciseName = 5
	LungeExerciseNameBarbellBoxLunge                               LungeExerciseName = 6
	LungeExerciseNameBarbellBulgarianSplitSquat                    LungeExerciseName = 7
	LungeExerciseNameBarbellCrossoverLunge                         LungeExerciseName = 8
	LungeExerciseNameBarbellFrontSplitSquat                        LungeExerciseName = 9
	LungeExerciseNameBarbellLunge                                  LungeExerciseName = 10
	LungeExerciseNameBarbellReverseLunge                           LungeExerciseName = 11
	LungeExerciseNameBarbellSideLunge                              LungeExerciseName = 12
	LungeExerciseNameBarbellSplitSquat                             LungeExerciseName = 13
	LungeExerciseNameCoreControlRearLunge                          LungeExerciseName = 14
	LungeExerciseNameDiagonalLunge                                 LungeExerciseName = 15
	LungeExerciseNameDropLunge                                     LungeExerciseName = 16
	LungeExerciseNameDumbbellBoxLunge                              LungeExerciseName = 17
	LungeExerciseNameDumbbellBulgarianSplitSquat                   LungeExerciseName = 18
	LungeExerciseNameDumbbellCrossoverLunge                        LungeExerciseName = 19
	LungeExerciseNameDumbbellDiagonalLunge                         LungeExerciseName = 20
	LungeExerciseNameDumbbellLunge                                 LungeExerciseName = 21
	LungeExerciseNameDumbbellLungeAndRotation                      LungeExerciseName = 22
	LungeExerciseNameDumbbellOverheadBulgarianSplitSquat           LungeExerciseName = 23
	LungeExerciseNameDumbbellReverseLungeToHighKneeAndPress        LungeExerciseName = 24
	LungeExerciseNameDumbbellSideLunge                             LungeExerciseName = 25
	LungeExerciseNameElevatedFrontFootBarbellSplitSquat            LungeExerciseName = 26
	LungeExerciseNameFrontFootElevatedDumbbellSplitSquat           LungeExerciseName = 27
	LungeExerciseNameGunslingerLunge                               LungeExerciseName = 28
	LungeExerciseNameLawnmowerLunge                                LungeExerciseName = 29
	LungeExerciseNameLowLungeWithIsometricAdduction                LungeExerciseName = 30
	LungeExerciseNameLowSideToSideLunge                            LungeExerciseName = 31
	LungeExerciseNameLunge                                         LungeExerciseName = 32
	LungeExerciseNameWeightedLunge                                 LungeExerciseName = 33
	LungeExerciseNameLungeWithArmReach                             LungeExerciseName = 34
	LungeExerciseNameLungeWithDiagonalReach                        LungeExerciseName = 35
	LungeExerciseNameLungeWithSideBend                             LungeExerciseName = 36
	LungeExerciseNameOffsetDumbbellLunge                           LungeExerciseName = 37
	LungeExerciseNameOffsetDumbbellReverseLunge                    LungeExerciseName = 38
	LungeExerciseNameOverheadBulgarianSplitSquat                   LungeExerciseName = 39
	LungeExerciseNameOverheadDumbbellReverseLunge                  LungeExerciseName = 40
	LungeExerciseNameOverheadDumbbellSplitSquat                    LungeExerciseName = 41
	LungeExerciseNameOverheadLungeWithRotation                     LungeExerciseName = 42
	LungeExerciseNameReverseBarbellBoxLunge                        LungeExerciseName = 43
	LungeExerciseNameReverseBoxLunge                               LungeExerciseName = 44
	LungeExerciseNameReverseDumbbellBoxLunge                       LungeExerciseName = 45
	LungeExerciseNameReverseDumbbellCrossoverLunge                 LungeExerciseName = 46
	LungeExerciseNameReverseDumbbellDiagonalLunge                  LungeExerciseName = 47
	LungeExerciseNameReverseLungeWithReachBack                     LungeExerciseName = 48
	LungeExerciseNameWeightedReverseLungeWithReachBack             LungeExerciseName = 49
	LungeExerciseNameReverseLungeWithTwistAndOverheadReach         LungeExerciseName = 50
	LungeExerciseNameWeightedReverseLungeWithTwistAndOverheadReach LungeExerciseName = 51
	LungeExerciseNameReverseSlidingBoxLunge                        LungeExerciseName = 52
	LungeExerciseNameWeightedReverseSlidingBoxLunge                LungeExerciseName = 53
	LungeExerciseNameReverseSlidingLunge                           LungeExerciseName = 54
	LungeExerciseNameWeightedReverseSlidingLunge                   LungeExerciseName = 55
	LungeExerciseNameRunnersLungeToBalance                         LungeExerciseName = 56
	LungeExerciseNameWeightedRunnersLungeToBalance                 LungeExerciseName = 57
	LungeExerciseNameShiftingSideLunge                             LungeExerciseName = 58
	LungeExerciseNameSideAndCrossoverLunge                         LungeExerciseName = 59
	LungeExerciseNameWeightedSideAndCrossoverLunge                 LungeExerciseName = 60
	LungeExerciseNameSideLunge                                     LungeExerciseName = 61
	LungeExerciseNameWeightedSideLunge                             LungeExerciseName = 62
	LungeExerciseNameSideLungeAndPress                             LungeExerciseName = 63
	LungeExerciseNameSideLungeJumpOff                              LungeExerciseName = 64
	LungeExerciseNameSideLungeSweep                                LungeExerciseName = 65
	LungeExerciseNameWeightedSideLungeSweep                        LungeExerciseName = 66
	LungeExerciseNameSideLungeToCrossoverTap                       LungeExerciseName = 67
	LungeExerciseNameWeightedSideLungeToCrossoverTap               LungeExerciseName = 68
	LungeExerciseNameSideToSideLungeChops                          LungeExerciseName = 69
	LungeExerciseNameWeightedSideToSideLungeChops                  LungeExerciseName = 70
	LungeExerciseNameSiffJumpLunge                                 LungeExerciseName = 71
	LungeExerciseNameWeightedSiffJumpLunge                         LungeExerciseName = 72
	LungeExerciseNameSingleArmReverseLungeAndPress                 LungeExerciseName = 73
	LungeExerciseNameSlidingLateralLunge                           LungeExerciseName = 74
	LungeExerciseNameWeightedSlidingLateralLunge                   LungeExerciseName = 75
	LungeExerciseNameWalkingBarbellLunge                           LungeExerciseName = 76
	LungeExerciseNameWalkingDumbbellLunge                          LungeExerciseName = 77
	LungeExerciseNameWalkingLunge                                  LungeExerciseName = 78
	LungeExerciseNameWeightedWalkingLunge                          LungeExerciseName = 79
	LungeExerciseNameWideGripOverheadBarbellSplitSquat             LungeExerciseName = 80
	LungeExerciseNameInvalid                                       LungeExerciseName = 0xFFFF
)

func (l LungeExerciseName) Uint16() uint16 { return uint16(l) }

func (l LungeExerciseName) String() string {
	switch l {
	case LungeExerciseNameOverheadLunge:
		return "overhead_lunge"
	case LungeExerciseNameLungeMatrix:
		return "lunge_matrix"
	case LungeExerciseNameWeightedLungeMatrix:
		return "weighted_lunge_matrix"
	case LungeExerciseNameAlternatingBarbellForwardLunge:
		return "alternating_barbell_forward_lunge"
	case LungeExerciseNameAlternatingDumbbellLungeWithReach:
		return "alternating_dumbbell_lunge_with_reach"
	case LungeExerciseNameBackFootElevatedDumbbellSplitSquat:
		return "back_foot_elevated_dumbbell_split_squat"
	case LungeExerciseNameBarbellBoxLunge:
		return "barbell_box_lunge"
	case LungeExerciseNameBarbellBulgarianSplitSquat:
		return "barbell_bulgarian_split_squat"
	case LungeExerciseNameBarbellCrossoverLunge:
		return "barbell_crossover_lunge"
	case LungeExerciseNameBarbellFrontSplitSquat:
		return "barbell_front_split_squat"
	case LungeExerciseNameBarbellLunge:
		return "barbell_lunge"
	case LungeExerciseNameBarbellReverseLunge:
		return "barbell_reverse_lunge"
	case LungeExerciseNameBarbellSideLunge:
		return "barbell_side_lunge"
	case LungeExerciseNameBarbellSplitSquat:
		return "barbell_split_squat"
	case LungeExerciseNameCoreControlRearLunge:
		return "core_control_rear_lunge"
	case LungeExerciseNameDiagonalLunge:
		return "diagonal_lunge"
	case LungeExerciseNameDropLunge:
		return "drop_lunge"
	case LungeExerciseNameDumbbellBoxLunge:
		return "dumbbell_box_lunge"
	case LungeExerciseNameDumbbellBulgarianSplitSquat:
		return "dumbbell_bulgarian_split_squat"
	case LungeExerciseNameDumbbellCrossoverLunge:
		return "dumbbell_crossover_lunge"
	case LungeExerciseNameDumbbellDiagonalLunge:
		return "dumbbell_diagonal_lunge"
	case LungeExerciseNameDumbbellLunge:
		return "dumbbell_lunge"
	case LungeExerciseNameDumbbellLungeAndRotation:
		return "dumbbell_lunge_and_rotation"
	case LungeExerciseNameDumbbellOverheadBulgarianSplitSquat:
		return "dumbbell_overhead_bulgarian_split_squat"
	case LungeExerciseNameDumbbellReverseLungeToHighKneeAndPress:
		return "dumbbell_reverse_lunge_to_high_knee_and_press"
	case LungeExerciseNameDumbbellSideLunge:
		return "dumbbell_side_lunge"
	case LungeExerciseNameElevatedFrontFootBarbellSplitSquat:
		return "elevated_front_foot_barbell_split_squat"
	case LungeExerciseNameFrontFootElevatedDumbbellSplitSquat:
		return "front_foot_elevated_dumbbell_split_squat"
	case LungeExerciseNameGunslingerLunge:
		return "gunslinger_lunge"
	case LungeExerciseNameLawnmowerLunge:
		return "lawnmower_lunge"
	case LungeExerciseNameLowLungeWithIsometricAdduction:
		return "low_lunge_with_isometric_adduction"
	case LungeExerciseNameLowSideToSideLunge:
		return "low_side_to_side_lunge"
	case LungeExerciseNameLunge:
		return "lunge"
	case LungeExerciseNameWeightedLunge:
		return "weighted_lunge"
	case LungeExerciseNameLungeWithArmReach:
		return "lunge_with_arm_reach"
	case LungeExerciseNameLungeWithDiagonalReach:
		return "lunge_with_diagonal_reach"
	case LungeExerciseNameLungeWithSideBend:
		return "lunge_with_side_bend"
	case LungeExerciseNameOffsetDumbbellLunge:
		return "offset_dumbbell_lunge"
	case LungeExerciseNameOffsetDumbbellReverseLunge:
		return "offset_dumbbell_reverse_lunge"
	case LungeExerciseNameOverheadBulgarianSplitSquat:
		return "overhead_bulgarian_split_squat"
	case LungeExerciseNameOverheadDumbbellReverseLunge:
		return "overhead_dumbbell_reverse_lunge"
	case LungeExerciseNameOverheadDumbbellSplitSquat:
		return "overhead_dumbbell_split_squat"
	case LungeExerciseNameOverheadLungeWithRotation:
		return "overhead_lunge_with_rotation"
	case LungeExerciseNameReverseBarbellBoxLunge:
		return "reverse_barbell_box_lunge"
	case LungeExerciseNameReverseBoxLunge:
		return "reverse_box_lunge"
	case LungeExerciseNameReverseDumbbellBoxLunge:
		return "reverse_dumbbell_box_lunge"
	case LungeExerciseNameReverseDumbbellCrossoverLunge:
		return "reverse_dumbbell_crossover_lunge"
	case LungeExerciseNameReverseDumbbellDiagonalLunge:
		return "reverse_dumbbell_diagonal_lunge"
	case LungeExerciseNameReverseLungeWithReachBack:
		return "reverse_lunge_with_reach_back"
	case LungeExerciseNameWeightedReverseLungeWithReachBack:
		return "weighted_reverse_lunge_with_reach_back"
	case LungeExerciseNameReverseLungeWithTwistAndOverheadReach:
		return "reverse_lunge_with_twist_and_overhead_reach"
	case LungeExerciseNameWeightedReverseLungeWithTwistAndOverheadReach:
		return "weighted_reverse_lunge_with_twist_and_overhead_reach"
	case LungeExerciseNameReverseSlidingBoxLunge:
		return "reverse_sliding_box_lunge"
	case LungeExerciseNameWeightedReverseSlidingBoxLunge:
		return "weighted_reverse_sliding_box_lunge"
	case LungeExerciseNameReverseSlidingLunge:
		return "reverse_sliding_lunge"
	case LungeExerciseNameWeightedReverseSlidingLunge:
		return "weighted_reverse_sliding_lunge"
	case LungeExerciseNameRunnersLungeToBalance:
		return "runners_lunge_to_balance"
	case LungeExerciseNameWeightedRunnersLungeToBalance:
		return "weighted_runners_lunge_to_balance"
	case LungeExerciseNameShiftingSideLunge:
		return "shifting_side_lunge"
	case LungeExerciseNameSideAndCrossoverLunge:
		return "side_and_crossover_lunge"
	case LungeExerciseNameWeightedSideAndCrossoverLunge:
		return "weighted_side_and_crossover_lunge"
	case LungeExerciseNameSideLunge:
		return "side_lunge"
	case LungeExerciseNameWeightedSideLunge:
		return "weighted_side_lunge"
	case LungeExerciseNameSideLungeAndPress:
		return "side_lunge_and_press"
	case LungeExerciseNameSideLungeJumpOff:
		return "side_lunge_jump_off"
	case LungeExerciseNameSideLungeSweep:
		return "side_lunge_sweep"
	case LungeExerciseNameWeightedSideLungeSweep:
		return "weighted_side_lunge_sweep"
	case LungeExerciseNameSideLungeToCrossoverTap:
		return "side_lunge_to_crossover_tap"
	case LungeExerciseNameWeightedSideLungeToCrossoverTap:
		return "weighted_side_lunge_to_crossover_tap"
	case LungeExerciseNameSideToSideLungeChops:
		return "side_to_side_lunge_chops"
	case LungeExerciseNameWeightedSideToSideLungeChops:
		return "weighted_side_to_side_lunge_chops"
	case LungeExerciseNameSiffJumpLunge:
		return "siff_jump_lunge"
	case LungeExerciseNameWeightedSiffJumpLunge:
		return "weighted_siff_jump_lunge"
	case LungeExerciseNameSingleArmReverseLungeAndPress:
		return "single_arm_reverse_lunge_and_press"
	case LungeExerciseNameSlidingLateralLunge:
		return "sliding_lateral_lunge"
	case LungeExerciseNameWeightedSlidingLateralLunge:
		return "weighted_sliding_lateral_lunge"
	case LungeExerciseNameWalkingBarbellLunge:
		return "walking_barbell_lunge"
	case LungeExerciseNameWalkingDumbbellLunge:
		return "walking_dumbbell_lunge"
	case LungeExerciseNameWalkingLunge:
		return "walking_lunge"
	case LungeExerciseNameWeightedWalkingLunge:
		return "weighted_walking_lunge"
	case LungeExerciseNameWideGripOverheadBarbellSplitSquat:
		return "wide_grip_overhead_barbell_split_squat"
	default:
		return "LungeExerciseNameInvalid(" + strconv.FormatUint(uint64(l), 10) + ")"
	}
}

// FromString parse string into LungeExerciseName constant it's represent, return LungeExerciseNameInvalid if not found.
func LungeExerciseNameFromString(s string) LungeExerciseName {
	switch s {
	case "overhead_lunge":
		return LungeExerciseNameOverheadLunge
	case "lunge_matrix":
		return LungeExerciseNameLungeMatrix
	case "weighted_lunge_matrix":
		return LungeExerciseNameWeightedLungeMatrix
	case "alternating_barbell_forward_lunge":
		return LungeExerciseNameAlternatingBarbellForwardLunge
	case "alternating_dumbbell_lunge_with_reach":
		return LungeExerciseNameAlternatingDumbbellLungeWithReach
	case "back_foot_elevated_dumbbell_split_squat":
		return LungeExerciseNameBackFootElevatedDumbbellSplitSquat
	case "barbell_box_lunge":
		return LungeExerciseNameBarbellBoxLunge
	case "barbell_bulgarian_split_squat":
		return LungeExerciseNameBarbellBulgarianSplitSquat
	case "barbell_crossover_lunge":
		return LungeExerciseNameBarbellCrossoverLunge
	case "barbell_front_split_squat":
		return LungeExerciseNameBarbellFrontSplitSquat
	case "barbell_lunge":
		return LungeExerciseNameBarbellLunge
	case "barbell_reverse_lunge":
		return LungeExerciseNameBarbellReverseLunge
	case "barbell_side_lunge":
		return LungeExerciseNameBarbellSideLunge
	case "barbell_split_squat":
		return LungeExerciseNameBarbellSplitSquat
	case "core_control_rear_lunge":
		return LungeExerciseNameCoreControlRearLunge
	case "diagonal_lunge":
		return LungeExerciseNameDiagonalLunge
	case "drop_lunge":
		return LungeExerciseNameDropLunge
	case "dumbbell_box_lunge":
		return LungeExerciseNameDumbbellBoxLunge
	case "dumbbell_bulgarian_split_squat":
		return LungeExerciseNameDumbbellBulgarianSplitSquat
	case "dumbbell_crossover_lunge":
		return LungeExerciseNameDumbbellCrossoverLunge
	case "dumbbell_diagonal_lunge":
		return LungeExerciseNameDumbbellDiagonalLunge
	case "dumbbell_lunge":
		return LungeExerciseNameDumbbellLunge
	case "dumbbell_lunge_and_rotation":
		return LungeExerciseNameDumbbellLungeAndRotation
	case "dumbbell_overhead_bulgarian_split_squat":
		return LungeExerciseNameDumbbellOverheadBulgarianSplitSquat
	case "dumbbell_reverse_lunge_to_high_knee_and_press":
		return LungeExerciseNameDumbbellReverseLungeToHighKneeAndPress
	case "dumbbell_side_lunge":
		return LungeExerciseNameDumbbellSideLunge
	case "elevated_front_foot_barbell_split_squat":
		return LungeExerciseNameElevatedFrontFootBarbellSplitSquat
	case "front_foot_elevated_dumbbell_split_squat":
		return LungeExerciseNameFrontFootElevatedDumbbellSplitSquat
	case "gunslinger_lunge":
		return LungeExerciseNameGunslingerLunge
	case "lawnmower_lunge":
		return LungeExerciseNameLawnmowerLunge
	case "low_lunge_with_isometric_adduction":
		return LungeExerciseNameLowLungeWithIsometricAdduction
	case "low_side_to_side_lunge":
		return LungeExerciseNameLowSideToSideLunge
	case "lunge":
		return LungeExerciseNameLunge
	case "weighted_lunge":
		return LungeExerciseNameWeightedLunge
	case "lunge_with_arm_reach":
		return LungeExerciseNameLungeWithArmReach
	case "lunge_with_diagonal_reach":
		return LungeExerciseNameLungeWithDiagonalReach
	case "lunge_with_side_bend":
		return LungeExerciseNameLungeWithSideBend
	case "offset_dumbbell_lunge":
		return LungeExerciseNameOffsetDumbbellLunge
	case "offset_dumbbell_reverse_lunge":
		return LungeExerciseNameOffsetDumbbellReverseLunge
	case "overhead_bulgarian_split_squat":
		return LungeExerciseNameOverheadBulgarianSplitSquat
	case "overhead_dumbbell_reverse_lunge":
		return LungeExerciseNameOverheadDumbbellReverseLunge
	case "overhead_dumbbell_split_squat":
		return LungeExerciseNameOverheadDumbbellSplitSquat
	case "overhead_lunge_with_rotation":
		return LungeExerciseNameOverheadLungeWithRotation
	case "reverse_barbell_box_lunge":
		return LungeExerciseNameReverseBarbellBoxLunge
	case "reverse_box_lunge":
		return LungeExerciseNameReverseBoxLunge
	case "reverse_dumbbell_box_lunge":
		return LungeExerciseNameReverseDumbbellBoxLunge
	case "reverse_dumbbell_crossover_lunge":
		return LungeExerciseNameReverseDumbbellCrossoverLunge
	case "reverse_dumbbell_diagonal_lunge":
		return LungeExerciseNameReverseDumbbellDiagonalLunge
	case "reverse_lunge_with_reach_back":
		return LungeExerciseNameReverseLungeWithReachBack
	case "weighted_reverse_lunge_with_reach_back":
		return LungeExerciseNameWeightedReverseLungeWithReachBack
	case "reverse_lunge_with_twist_and_overhead_reach":
		return LungeExerciseNameReverseLungeWithTwistAndOverheadReach
	case "weighted_reverse_lunge_with_twist_and_overhead_reach":
		return LungeExerciseNameWeightedReverseLungeWithTwistAndOverheadReach
	case "reverse_sliding_box_lunge":
		return LungeExerciseNameReverseSlidingBoxLunge
	case "weighted_reverse_sliding_box_lunge":
		return LungeExerciseNameWeightedReverseSlidingBoxLunge
	case "reverse_sliding_lunge":
		return LungeExerciseNameReverseSlidingLunge
	case "weighted_reverse_sliding_lunge":
		return LungeExerciseNameWeightedReverseSlidingLunge
	case "runners_lunge_to_balance":
		return LungeExerciseNameRunnersLungeToBalance
	case "weighted_runners_lunge_to_balance":
		return LungeExerciseNameWeightedRunnersLungeToBalance
	case "shifting_side_lunge":
		return LungeExerciseNameShiftingSideLunge
	case "side_and_crossover_lunge":
		return LungeExerciseNameSideAndCrossoverLunge
	case "weighted_side_and_crossover_lunge":
		return LungeExerciseNameWeightedSideAndCrossoverLunge
	case "side_lunge":
		return LungeExerciseNameSideLunge
	case "weighted_side_lunge":
		return LungeExerciseNameWeightedSideLunge
	case "side_lunge_and_press":
		return LungeExerciseNameSideLungeAndPress
	case "side_lunge_jump_off":
		return LungeExerciseNameSideLungeJumpOff
	case "side_lunge_sweep":
		return LungeExerciseNameSideLungeSweep
	case "weighted_side_lunge_sweep":
		return LungeExerciseNameWeightedSideLungeSweep
	case "side_lunge_to_crossover_tap":
		return LungeExerciseNameSideLungeToCrossoverTap
	case "weighted_side_lunge_to_crossover_tap":
		return LungeExerciseNameWeightedSideLungeToCrossoverTap
	case "side_to_side_lunge_chops":
		return LungeExerciseNameSideToSideLungeChops
	case "weighted_side_to_side_lunge_chops":
		return LungeExerciseNameWeightedSideToSideLungeChops
	case "siff_jump_lunge":
		return LungeExerciseNameSiffJumpLunge
	case "weighted_siff_jump_lunge":
		return LungeExerciseNameWeightedSiffJumpLunge
	case "single_arm_reverse_lunge_and_press":
		return LungeExerciseNameSingleArmReverseLungeAndPress
	case "sliding_lateral_lunge":
		return LungeExerciseNameSlidingLateralLunge
	case "weighted_sliding_lateral_lunge":
		return LungeExerciseNameWeightedSlidingLateralLunge
	case "walking_barbell_lunge":
		return LungeExerciseNameWalkingBarbellLunge
	case "walking_dumbbell_lunge":
		return LungeExerciseNameWalkingDumbbellLunge
	case "walking_lunge":
		return LungeExerciseNameWalkingLunge
	case "weighted_walking_lunge":
		return LungeExerciseNameWeightedWalkingLunge
	case "wide_grip_overhead_barbell_split_squat":
		return LungeExerciseNameWideGripOverheadBarbellSplitSquat
	default:
		return LungeExerciseNameInvalid
	}
}

// List returns all constants.
func ListLungeExerciseName() []LungeExerciseName {
	return []LungeExerciseName{
		LungeExerciseNameOverheadLunge,
		LungeExerciseNameLungeMatrix,
		LungeExerciseNameWeightedLungeMatrix,
		LungeExerciseNameAlternatingBarbellForwardLunge,
		LungeExerciseNameAlternatingDumbbellLungeWithReach,
		LungeExerciseNameBackFootElevatedDumbbellSplitSquat,
		LungeExerciseNameBarbellBoxLunge,
		LungeExerciseNameBarbellBulgarianSplitSquat,
		LungeExerciseNameBarbellCrossoverLunge,
		LungeExerciseNameBarbellFrontSplitSquat,
		LungeExerciseNameBarbellLunge,
		LungeExerciseNameBarbellReverseLunge,
		LungeExerciseNameBarbellSideLunge,
		LungeExerciseNameBarbellSplitSquat,
		LungeExerciseNameCoreControlRearLunge,
		LungeExerciseNameDiagonalLunge,
		LungeExerciseNameDropLunge,
		LungeExerciseNameDumbbellBoxLunge,
		LungeExerciseNameDumbbellBulgarianSplitSquat,
		LungeExerciseNameDumbbellCrossoverLunge,
		LungeExerciseNameDumbbellDiagonalLunge,
		LungeExerciseNameDumbbellLunge,
		LungeExerciseNameDumbbellLungeAndRotation,
		LungeExerciseNameDumbbellOverheadBulgarianSplitSquat,
		LungeExerciseNameDumbbellReverseLungeToHighKneeAndPress,
		LungeExerciseNameDumbbellSideLunge,
		LungeExerciseNameElevatedFrontFootBarbellSplitSquat,
		LungeExerciseNameFrontFootElevatedDumbbellSplitSquat,
		LungeExerciseNameGunslingerLunge,
		LungeExerciseNameLawnmowerLunge,
		LungeExerciseNameLowLungeWithIsometricAdduction,
		LungeExerciseNameLowSideToSideLunge,
		LungeExerciseNameLunge,
		LungeExerciseNameWeightedLunge,
		LungeExerciseNameLungeWithArmReach,
		LungeExerciseNameLungeWithDiagonalReach,
		LungeExerciseNameLungeWithSideBend,
		LungeExerciseNameOffsetDumbbellLunge,
		LungeExerciseNameOffsetDumbbellReverseLunge,
		LungeExerciseNameOverheadBulgarianSplitSquat,
		LungeExerciseNameOverheadDumbbellReverseLunge,
		LungeExerciseNameOverheadDumbbellSplitSquat,
		LungeExerciseNameOverheadLungeWithRotation,
		LungeExerciseNameReverseBarbellBoxLunge,
		LungeExerciseNameReverseBoxLunge,
		LungeExerciseNameReverseDumbbellBoxLunge,
		LungeExerciseNameReverseDumbbellCrossoverLunge,
		LungeExerciseNameReverseDumbbellDiagonalLunge,
		LungeExerciseNameReverseLungeWithReachBack,
		LungeExerciseNameWeightedReverseLungeWithReachBack,
		LungeExerciseNameReverseLungeWithTwistAndOverheadReach,
		LungeExerciseNameWeightedReverseLungeWithTwistAndOverheadReach,
		LungeExerciseNameReverseSlidingBoxLunge,
		LungeExerciseNameWeightedReverseSlidingBoxLunge,
		LungeExerciseNameReverseSlidingLunge,
		LungeExerciseNameWeightedReverseSlidingLunge,
		LungeExerciseNameRunnersLungeToBalance,
		LungeExerciseNameWeightedRunnersLungeToBalance,
		LungeExerciseNameShiftingSideLunge,
		LungeExerciseNameSideAndCrossoverLunge,
		LungeExerciseNameWeightedSideAndCrossoverLunge,
		LungeExerciseNameSideLunge,
		LungeExerciseNameWeightedSideLunge,
		LungeExerciseNameSideLungeAndPress,
		LungeExerciseNameSideLungeJumpOff,
		LungeExerciseNameSideLungeSweep,
		LungeExerciseNameWeightedSideLungeSweep,
		LungeExerciseNameSideLungeToCrossoverTap,
		LungeExerciseNameWeightedSideLungeToCrossoverTap,
		LungeExerciseNameSideToSideLungeChops,
		LungeExerciseNameWeightedSideToSideLungeChops,
		LungeExerciseNameSiffJumpLunge,
		LungeExerciseNameWeightedSiffJumpLunge,
		LungeExerciseNameSingleArmReverseLungeAndPress,
		LungeExerciseNameSlidingLateralLunge,
		LungeExerciseNameWeightedSlidingLateralLunge,
		LungeExerciseNameWalkingBarbellLunge,
		LungeExerciseNameWalkingDumbbellLunge,
		LungeExerciseNameWalkingLunge,
		LungeExerciseNameWeightedWalkingLunge,
		LungeExerciseNameWideGripOverheadBarbellSplitSquat,
	}
}