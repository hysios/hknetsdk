package netsdk

type IllegalType int

const (
	IllNormal            IllegalType = iota // 0-正常
	IllLowspeed                             // 1-低速
	IllOverspeed                            // 2-超速，
	IllRetrograde                           // 3-逆行，
	IllRedlight                             // 4-闯红灯
	IllOverLaneline                         // 5-压车道线,
	IllNoFollowGuided                       // 6-不按导向
	IllStuckIntersection                    // 7-路口滞留，
	IllOverNonvehicle                       // 8-机占非
	IllLaneChange                           // 9-违法变道
	IllNotByLane                            // 10-不按车道
	IllBreakBan                             // 11-违反禁令
	IllIntersectionStop                     // 12-路口停车
	IllStopGreen                            // 13-绿灯停车
	IllPeopleFirst                          // 14-未礼让行人(违法代码1357),
	IllTrafficParking                       // 15-违章停车
	IllTurn                                 // 16-违章掉头
	IllOccupyEmergency                      // 17-占用应急车道,
	IllBanRight                             // 18-禁右
	IllBanLeft                              // 19-禁左
	IllOverYellow                           // 20-压黄线
	IllUnwearingSeatbelt                    // 21-未系安全带
	IllPeopleRedlight                       // 22-行人闯红灯
	IllCarInsertion                         // 23-加塞
	IllHighBeam                             // 24-违法使用远光灯，
	IllCalling                              // 25-驾驶时拨打接听手持电话
	IllNotYieldLeft                         // 26-左转不让直行，
	IllNotYieldRight                        // 27-右转不让左转，
	IllNotYieldTurn                         // 28-掉头不让直行，
	IllTurnSmall                            // 29-大弯小转,
	IllBreakGreen                           // 30-闯绿灯，
	IllUnhelmet                             // 31-未带头盔，
	IllCarryMan                             // 32-非机动车载人，
	IllOccupyVehicleLane                    // 33-非机动车占用机动车道，
	IllUmbrellaShed                         // 34-非机动车打伞棚,
	IllSmokyCar                             // 35-黑烟车,
	IllHonking                              // 36-鸣笛,
	IllStopOverLine                         // 37-压线停车,
	IllCrossLineStop                        // 38-跨位停车,
	IllOverAndCrossLine                     // 39-压线且跨位停车
	// Ill 40-不让右方道路来车先行 41-进入环形路口未让已在路口内的机动车先行 42-机动车从匝道进入主路未让行
)
