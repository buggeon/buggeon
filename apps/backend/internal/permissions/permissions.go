package permissions

type Role string

const (
	RoleOwner   Role = "owner"
	RoleManager Role = "manager"
	RoleMember  Role = "member"
)

var RolePermissions = map[Role][]string{
	RoleOwner: {
		"project.get",
		"project.delete",
		"project.update",
	},
}
