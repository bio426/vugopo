<script setup lang="ts">
import { reactive, ref } from "vue"
import { useRouter } from "vue-router"

import taskService from "@/service/task"

const router = useRouter()
const loading = ref(false)
const form = reactive({ title: "", public: false })

async function create() {
	loading.value = true
	await taskService.create({ title: form.title, public: form.public })
	loading.value = false
	router.push({ name: "task" })
}
</script>

<template>
	<div>
		<h4>Create Task</h4>
		<form @submit.prevent="create">
			<label>
				Title
				<input type="text" required v-model="form.title" />
			</label>
			<fieldset>
				<label>
					<input type="checkbox" v-model="form.public" />
					Public
				</label>
			</fieldset>
			<button :aria-busy="loading" type="submit">Create</button>
		</form>
	</div>
</template>
