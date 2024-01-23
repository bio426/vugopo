type UserRole = "super" | "admin" | "user"

export type User = {
	username: string
	role: UserRole
}

export type ExpirableUser = User & { expiryDate: string }
