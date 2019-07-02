package modules

// ConvertTstatusToMsg : convert vat status to message
func ConvertTstatusToMsg(tstatus string) string {
	var message string
	switch tstatus {
	case "initial":
	case "pending":
		message = "Notification Payment Pending"
	case "expire":
		message = "Notification Payment Failed"
	case "deny":
		message = "Notification Payment Failed"
	case "refund":
		message = "Notification Payment Failed"
	case "failure":
		message = "Notification Payment Failed"
	case "chargeback":
		message = "Notification Payment Failed"
	case "cancel":
		message = "Notification Payment Failed"
	case "success":
	case "settlement":
		message = "Notification Payment Received"
	default:
		message = "Notification Payment"
	}

	return message
}
