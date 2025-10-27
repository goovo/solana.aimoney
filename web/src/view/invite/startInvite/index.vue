<template>
  <div>
    <div class="gva-table-box">
      <!-- 推荐信息卡片 -->
      <el-row :gutter="20">
        <el-col :span="8">
          <el-card shadow="hover" class="invite-card">
            <template #header>
              <div class="card-header">
                <span>我的推荐码</span>
                <el-tag type="success">专属</el-tag>
              </div>
            </template>
            <div class="invite-content">
              <div class="invite-code">
                <span class="code-text">{{ inviteInfo.inviteCode || '加载中...' }}</span>
                <el-button 
                  type="primary" 
                  link 
                  @click="copyText(inviteInfo.inviteCode)"
                  :icon="CopyDocument"
                >
                  复制
                </el-button>
              </div>
            </div>
          </el-card>
        </el-col>
        
        <el-col :span="8">
          <el-card shadow="hover" class="invite-card">
            <template #header>
              <div class="card-header">
                <span>推荐链接</span>
                <el-tag type="warning">分享</el-tag>
              </div>
            </template>
            <div class="invite-content">
              <div class="invite-link">
                <el-input 
                  v-model="inviteInfo.inviteLink" 
                  readonly 
                  placeholder="加载中..."
                >
                  <template #append>
                    <el-button @click="copyText(inviteInfo.inviteLink)" :icon="CopyDocument" />
                  </template>
                </el-input>
              </div>
            </div>
          </el-card>
        </el-col>
        
        <el-col :span="8">
          <el-card shadow="hover" class="invite-card">
            <template #header>
              <div class="card-header">
                <span>推荐人数</span>
                <el-tag type="info">统计</el-tag>
              </div>
            </template>
            <div class="invite-content">
              <div class="referral-count">
                <span class="count-number">{{ inviteInfo.totalReferral || 0 }}</span>
                <span class="count-text">人</span>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>

      <!-- 推荐说明 -->
      <el-card shadow="never" class="invite-desc-card">
        <template #header>
          <span>推荐说明</span>
        </template>
        <div class="invite-desc">
          <el-alert
            title="如何获得推荐奖励？"
            type="info"
            :closable="false"
            show-icon
          >
            <div>
              <p>1. 分享您的专属推荐链接或推荐码给您的朋友</p>
              <p>2. 您的朋友通过链接注册或填写您的推荐码</p>
              <p>3. 您的朋友完成注册后，成功投资30天后，系统将计算您的收益</p>
            </div>
          </el-alert>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script setup>
defineOptions({
  name: 'StartInvite'
})

import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { CopyDocument } from '@element-plus/icons-vue'
import { getInviteInfo } from '@/api/invite'

const inviteInfo = ref({
  inviteCode: '',
  inviteLink: '',
  totalReferral: 0
})

// 获取推荐信息
const fetchInviteInfo = async () => {
  try {
    const res = await getInviteInfo()
    if (res.code === 0) {
      inviteInfo.value = res.data
    }
  } catch (error) {
    ElMessage.error('获取推荐信息失败')
  }
}

// 复制文本
const copyText = (text) => {
  if (!text) return
  
  navigator.clipboard.writeText(text).then(() => {
    ElMessage.success('复制成功')
  }).catch(() => {
    // 降级方案
    const textarea = document.createElement('textarea')
    textarea.value = text
    document.body.appendChild(textarea)
    textarea.select()
    document.execCommand('copy')
    document.body.removeChild(textarea)
    ElMessage.success('复制成功')
  })
}

onMounted(() => {
  fetchInviteInfo()
})
</script>

<style scoped lang="scss">
.invite-card {
  height: 180px;
  
  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-weight: bold;
  }
  
  .invite-content {
    text-align: center;
    padding: 20px 0;
    
    .invite-code {
      display: flex;
      align-items: center;
      justify-content: center;
      gap: 15px;
      
      .code-text {
        font-size: 24px;
        font-weight: bold;
        color: #409EFF;
        font-family: 'Courier New', monospace;
      }
    }
    
    .invite-link {
      width: 100%;
    }
    
    .referral-count {
      display: flex;
      align-items: baseline;
      justify-content: center;
      gap: 5px;
      
      .count-number {
        font-size: 36px;
        font-weight: bold;
        color: #67C23A;
      }
      
      .count-text {
        font-size: 14px;
        color: #909399;
      }
    }
  }
}

.invite-desc-card {
  margin-top: 20px;
  
  .invite-desc {
    padding: 10px 0;
    
    p {
      margin: 8px 0;
      line-height: 1.6;
    }
  }
}
</style>