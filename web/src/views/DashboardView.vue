<template>
  <Suspense>
    <template v-slot:default>
      <SmallStatics></SmallStatics>
    </template>

    <template v-slot:fallback>
      <el-skeleton :rows="4" animated translate="yes" />
    </template>
</Suspense>


  
</template>


<script setup lang="ts">

import {defineAsyncComponent} from 'vue'

const SmallStatics = defineAsyncComponent(() => import('../components/dashboard/SmallStatistics.vue'))


</script>
<style scoped>


</style>