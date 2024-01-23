<script setup lang="ts">
import { ref, onUnmounted } from "vue"

import * as tTask from "@/type/task"
import taskService from "@/service/task"

const loading = ref(false)

const rows = ref<tTask.PublicTask[]>([])

async function getRows() {
	loading.value = true
	const res = await taskService.getPublicTasks()
	rows.value = res.data
	loading.value = false
}
// getRows()

const events = new EventSource("http://localhost:1323/event/test")
events.addEventListener("message", (e: MessageEvent) => {
	console.log(e)
	console.log(e.data, e.type)
})

events.addEventListener("error", (e) => {
	console.log(e)
})

onUnmounted(() => events.close())
</script>

<template>
	<div>
		<h4>Popular Pubic Tasks</h4>
		<figure>
			<table>
				<thead>
					<tr>
						<th scope="col">Title</th>
						<th scope="col">Interested</th>
						<th scope="col">Done</th>
						<th scope="col">Created</th>
						<th scope="col">Owner</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="row in rows" :key="row.id">
						<td>{{ row.title }}</td>
						<td>
							<label>
								<input type="checkbox" disabled />
								{{ row.interesteds }}
							</label>
						</td>
						<td>
							<input
								type="checkbox"
								disabled
								v-model="row.done"
							/>
						</td>
						<td>
							{{ new Date(row.createdAt).toLocaleDateString() }}
						</td>
						<td>{{ row.owner }}</td>
					</tr>
					<tr v-if="rows.length == 0">
						<td>No data available</td>
					</tr>
				</tbody>
			</table>
		</figure>
	</div>
</template>
