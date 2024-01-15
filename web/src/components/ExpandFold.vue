<template>
  <div class="nav-header">
    <el-icon :size="30">
      <component class="fold-menu" :is="iconName" @click="handleFoldClick"></component>
    </el-icon>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  emits: ['foldChange'], //子传父
  setup(props, { emit }) {
    let iconName = ref('Fold')
    const isFold = ref(false)
    const handleFoldClick = () => {
      isFold.value = !isFold.value
      if (isFold.value == false) {
        iconName.value = 'Expand'
      } else {
        iconName.value = 'Fold'
      }
      emit('foldChange', isFold.value)
    }
    return {
      iconName,
      handleFoldClick,
      isFold
    }
  }
})
</script>

<style scoped>
.fold-menu {
  cursor: pointer;
}
</style>
