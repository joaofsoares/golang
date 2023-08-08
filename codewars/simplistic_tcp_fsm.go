package codewars

// there is an issue in this kata
// [APP_PASSIVE_OPEN RCV_SYN RCV_ACK APP_CLOSE APP_SEND]
// APP_CLOSE returns CLOSED, FIN_WAIT_1 and LAST_ACK, and APP_SEND is not available in any of those options
// it should return ERROR instead of FIN_WAIT_1

func TraverseTCPStates(events []string) string {
	states := map[string]map[string]string{
		"CLOSED": {
			"APP_PASSIVE_OPEN": "LISTEN",
			"APP_ACTIVE_OPEN":  "SYN_SENT",
		},
		"LISTEN": {
			"RCV_SYN":   "SYN_RCVD",
			"APP_SEND":  "SYN_SENT",
			"APP_CLOSE": "CLOSED",
		},
		"SYN_RCVD": {
			"APP_CLOSE": "FIN_WAIT_1",
			"RCV_ACK":   "ESTABLISHED",
		},
		"SYN_SENT": {
			"RCV_SYN":     "SYN_RCVD",
			"RCV_SYN_ACK": "ESTABLISHED",
			"APP_CLOSE":   "CLOSED",
		},
		"ESTABLISHED": {
			"APP_CLOSE": "FIN_WAIT_1",
			"RCV_FIN":   "CLOSE_WAIT",
		},
		"FIN_WAIT_1": {
			"RCV_FIN":     "CLOSING",
			"RCV_FIN_ACK": "TIME_WAIT",
			"RCV_ACK":     "FIN_WAIT_2",
		},
		"CLOSING": {
			"RCV_ACK": "TIME_WAIT",
		},
		"FIN_WAIT_2": {
			"RCV_FIN": "TIME_WAIT",
		},
		"TIME_WAIT": {
			"APP_TIMEOUT": "CLOSED",
		},
		"CLOSE_WAIT": {
			"APP_CLOSE": "LAST_ACK",
		},
		"LAST_ACK": {
			"RCV_ACK": "CLOSED",
		},
	}

	nextState := ""

	for cnt := 0; cnt < len(events); cnt++ {
		if cnt == 0 {
			state := getFirstEvent(events[cnt], states)

			if state == "" {
				return "ERROR"
			}

			nextState = state

			continue
		}

		state := events[cnt]

		value, exist := states[nextState][state]
		if exist {
			nextState = value
		} else {
			return "ERROR"
		}

	}

	return nextState
}

func getFirstEvent(event string, states map[string]map[string]string) string {
	for _, ps := range states {
		value, exist := ps[event]

		if exist {
			return value
		}
	}
	return ""
}
