package datamodels

type MConfig struct {
	XPich float64 `json:"xPich"`
	YPich float64 `json:"yPich"`
	ZPich float64 `json:"zPich"`
}

type NC struct {
	Config   MConfig `json:"config"`   // machine config like Axis Pich's and so
	U        float64 `json:"u"`        // u parameter
	RefPos   Point   `json:"refPos"`   // point wich is target at u
	Pos      Point   `json:"pos"`      // point wich machine is at ,
	Velocity Vector  `json:"velocity"` //velocity calculated by master,and do by sub threads
	//U        float64 `json:"u"`        // parameter u of which is passing
}
