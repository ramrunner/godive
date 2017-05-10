package planner

import (
	"fmt"
)

const (
	NUMCOMPS         = 16 //number of tissue compartments for VPM
	DELTAGASFRACTION = 0.00001
)

type compvals []float64 //compartment values

//context for the planner
type InputData struct {
	AdjustedCriticalRadiusHe   compvals
	AdjustedCriticalRadiusNe   compvals
	AltitudeDiveAlgorithmOff   bool
	AltitudeOfDive             float64
	AmbPressureOnsetOfImperm   compvals
	BarometricPressure         float64
	ConstantPressureOtherGases float64
	CritVolumeParameterLambda  float64
	CriticalRadiusHeMicrons    float64
	CriticalRadiusN2Microns    float64
	CriticalVolumeAlgorithmOff bool
	GasTensionOnsetOfImperm    compvals
	GradientOnsetOfImpermAtm   float64
	HeliumPressure             compvals
	HeliumTimeConstant         compvals
	InitialCriticalRadiusHe    compvals
	InitialCriticalRadiusN2    compvals
	MaxActualGradient          compvals
	MaxCrushingPressureHe      compvals
	MaxCrushingPressureN2      compvals
	MinimumDecoStopTime        float64
	NitrogenPressure           compvals
	NitrogenTimeConstant       compvals
	RegenerationTimeConstant   float64
	RunTime                    float64
	SegmentNumber              int
	SkinCompressionGammaC      float64
	SurfacePhaseVolumeTime     compvals
	SurfaceTensionGamma        float64
	UnitsFactor                float64
	UnitsWord1                 string
	UnitsWord2                 string
	WaterVaporPressure         float64
}

type Gas struct {
	FO2 float64
	FHe float64
	FN2 float64
}

func (g Gas) Assert() error {
	if (g.FO2+g.FHe+g.FN2)-1 > DELTAGASFRACTION {
		return fmt.Errorf("gas fractions don't add up to 1")
	}
	return nil
}

type DiveContext struct {
	Desc       string
	numGases   uint
	GasesUsed  []Gas
	AscentRate float64
}

func (d DiveContext) Assert() error {
	for _, g := range d.GasesUsed {
		if err := g.Assert(); err != nil {
			return fmt.Errorf("in gas :%+v :%s", g, err)
		}
	}
	return nil
}
