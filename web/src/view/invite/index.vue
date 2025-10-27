<template>
  <div>
    <div class="gva-btn-list">
      <el-tabs v-model="activeTab" type="border-card" @tab-click="handleTabClick">
        <el-tab-pane label="开启推荐" name="startInvite">
          <start-invite v-if="activeTab === 'startInvite'" />
        </el-tab-pane>
        <el-tab-pane label="我的推荐" name="myReferrals">
          <my-referrals v-if="activeTab === 'myReferrals'" />
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script setup>
defineOptions({
  name: 'Invite'
})

import { ref } from 'vue'
import StartInvite from './startInvite/index.vue'
import MyReferrals from './myReferrals/index.vue'

const activeTab = ref('startInvite')

const handleTabClick = (tab) => {
  activeTab.value = tab.props.name
}
</script>

<style scoped lang="scss">
.gva-btn-list {
  padding: 20px;
  background: #fff;
  border-radius: 4px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  
  .el-tabs {
    border: none;
    box-shadow: none;
  }
}
</style>