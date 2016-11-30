package planner

const (
	NUMCOMPS = 16 //number of tissue compartments for VPM
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
