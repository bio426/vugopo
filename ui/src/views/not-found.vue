<script setup lang="ts">
import { } from "vue"
</script>

<template>
    <div class="relative w-screen min-h-screen ">
        <div class="w-11/12 mx-auto">
            <!-- <Header /> -->
            <h1 class="my-8 text-2xl text-center font-bold">Not Found</h1>
            <div class="text-center">
                <router-link class="btn btn-link" :to="{ name: 'dashboard' }">To Dashboard</router-link>
            </div>
        </div>
    </div>
</template>
