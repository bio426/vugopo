<script setup lang="ts">
import { ref, onUnmounted } from "vue"

import * as tTask from "@/type/task"
import taskService from "@/service/task"
import Header from "@/components/Header.vue"

const loading = ref(false)

const rows = ref<tTask.PublicTask[]>([])

async function getRows() {
	loading.value = true
	const res = await taskService.getPublicTasks()
	rows.value = res.data
	loading.value = false
}
// getRows()

// const events = new EventSource("http://localhost:1323/event/test")
// events.addEventListener("message", (e: MessageEvent) => {
// 	console.log(e)
// 	console.log(e.data, e.type)
// })
//
// events.addEventListener("error", (e) => {
// 	console.log(e)
// })
// onUnmounted(() => events.close())
</script>

<template>
	<div class="relative w-screen min-h-screen">
		<div class="w-11/12 mx-auto">
			<Header title="Dashboard" />
			<p>
				Lorem ipsum, dolor sit amet consectetur adipisicing elit.
				Consequatur sed adipisci nihil vero amet aut, perferendis
				voluptate? Dignissimos corrupti ad qui commodi mollitia
				perspiciatis non, repellat fugiat error sunt vel!
			</p>
		</div>
	</div>
</template>
