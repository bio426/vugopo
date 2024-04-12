<script setup lang="ts">
import { ref } from "vue"

import * as tUser from "@/type/user"

import userService from "@/service/user"
import Header from "@/components/Header.vue"

const loading = ref(false)

const rows = ref<tUser.User[]>([])

async function getRows() {
    loading.value = true
    const res = await userService.list()
    rows.value = res.data
    loading.value = false
}
getRows()
</script>

<template>
    <div class="relative w-screen min-h-screen ">
        <div class="w-11/12 mx-auto">
            <Header title="Users" />
            <div class="mb-4">
                <router-link class="btn btn-primary" :to="{ name: 'user-create' }">Create</router-link>
            </div>
            <div class="overflow-x-auto">
                <table class="table">
                    <thead>
                        <tr>
                            <th>Username</th>
                            <th>Role</th>
                            <th>Created</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="row in rows" :key="row.id">
                            <td>{{ row.username }}</td>
                            <td>{{ row.role }}</td>
                            <td>{{ new Date(row.createdAt).toLocaleDateString() }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>
