import base from "./base"

import * as tTask from "@/type/task"

const path = "task"

export default {
	async getPublicTasks() {
		const res = await base.get(path + "/public")

		return res.json<{ data: tTask.PublicTask[] }>()
	},
	async getUserTask() {
		const res = await base.get(path)

		return res.json<{ data: tTask.Task[] }>()
	},
	async create(body: { title: string; public: boolean }) {
		const res = await base.post(path, { json: body })

		return res.status
	},
}
