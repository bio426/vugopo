export type Task = {
	id: number
	title: string
	public: boolean
	done: boolean
	createdAt: string
}

export type PublicTask = {
	id: number
	title: string
	done: boolean
	interesteds: number
	owner: string
	createdAt: string
}
