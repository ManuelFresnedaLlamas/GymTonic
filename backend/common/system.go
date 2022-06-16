package common

const TheSystem = "5c9ffa6b-ef90-4c8f-b818-b2ec3de87d2b"

func SystemContext(actionID string, role RoleType) AppContext {
	return AppContext{
		UserID:   TheSystem,
		ActionID: actionID,
		Role:     role,
	}
}
