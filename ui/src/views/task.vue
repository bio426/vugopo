<script setup lang="ts">
import { ref } from "vue"

import * as tTask from "@/type/task"
import taskService from "@/service/task"

const loading = ref(false)

const rows = ref<tTask.Task[]>([])

async function getRows() {
	loading.value = true
	const res = await taskService.getUserTask()
	rows.value = res.data
	loading.value = false
}
getRows()

// const evt = new EventSource()
// evt.close

if (window.Notification) {
	window.Notification.requestPermission()
}

function testNotif() {
	if (window.Notification.permission != "granted") return
	const imgSrc = new URL("../assets/images/misato.webp", import.meta.url).href
	const notif = new Notification("Test Notif", {
		body: "Some arbitrary text",
		icon: "/filter.png",
	})
	setTimeout(() => notif.close(), 2000)
}
</script>

<template>
	<div>
		<h4>My Tasks</h4>
		<section>
			<router-link role="button" :to="{ name: 'task-create' }">
				Create Task
			</router-link>
		</section>
		<figure :aria-busy="loading">
			<table>
				<thead>
					<tr>
						<th scope="col">Title</th>
						<th scope="col">Public</th>
						<th scope="col">Done</th>
						<th scope="col">Created</th>
						<th scope="col">Edit</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="row in rows" :key="row.id">
						<td>{{ row.title }}</td>
						<td>
							<input type="checkbox" v-model="row.public" />
						</td>
						<td>
							<input type="checkbox" v-model="row.done" />
						</td>
						<td>
							{{ new Date(row.createdAt).toLocaleDateString() }}
						</td>
						<td>
							<button>✍</button>
						</td>
					</tr>
					<tr v-if="rows.length == 0">
						<td colspan="5">No data available</td>
					</tr>
				</tbody>
			</table>
		</figure>
		<button @click="testNotif">Notif test</button>
	</div>
</template>
