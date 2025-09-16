<template>
  <div v-if="visible" class="modal-overlay">
    <!-- 左侧操作弹窗 -->
    <div v-if="showActions && cardActions?.length" class="card-actions">
      <button 
        v-for="(action, index) in cardActions" 
        :key="index" 
        @click="action.handler(card)"
      >
        {{ action.label }}
      </button>
    </div>

    <!-- 中间卡牌信息 -->
    <div class="card-modal">
      <img :src="card.avatar" class="card-avatar" />
      <div class="card-info">
        <h3 class="card-name">{{ card.name || '未知卡牌' }}</h3>
        <p class="card-rarity" :class="rarityClass">{{ card.rarity || '未知' }}</p>
        <div class="card-stats">
          <span>消耗: {{ card.cost || 0 }}</span>
          <span>攻击: {{ card.attack || 0 }}</span>
          <span>血量: {{ card.blood || 0 }}</span>
        </div>
        <p class="card-hashid">{{ card.hashid || '未知ID' }}</p>
        <button class="close-btn" @click="$emit('close')">关闭</button>
      </div>
    </div>

    <!-- 右侧卡牌描述 -->
    <div class="card-profile">
      <h4>卡牌描述</h4>
      <div class="profile-text">
        {{ card.profile || '暂无描述' }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "CardModal",
  props: {
    visible: Boolean,
    card: { type: Object, default: () => ({}) },
    showActions: { type: Boolean, default: false },
    cardActions: { type: Array, default: () => [] } // 左侧按钮数组，动态传入
  },
  computed: {
    rarityClass() {
      // 安全处理，避免 toLowerCase 报错
      const rarity = this.card?.rarity?.toLowerCase() || ''
      return rarity
    }
  }
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  z-index: 1000;
}

/* 左侧操作按钮弹窗 */
.card-actions {
  width: 140px;
  background: rgba(255, 255, 255, 0.9);
  border-radius: 6px;
  padding: 0.5rem;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  height: fit-content;
}

.card-actions button {
  padding: 0.4rem;
  border: none;
  border-radius: 4px;
  background: #42b983;
  color: #fff;
  cursor: pointer;
  text-align: center;
}

/* 中间卡牌信息 */
.card-modal {
  background: #fff;
  border-radius: 8px;
  width: 300px;
  padding: 1rem;
  text-align: left;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.card-avatar {
  width: 100%;
  border-radius: 6px;
  margin-bottom: 0.5rem;
}

.card-info {
  width: 100%;
}

.card-name {
  font-size: 1.2rem;
  margin-bottom: 0.3rem;
}

.card-rarity {
  font-weight: bold;
  margin-bottom: 0.5rem;
}

.card-rarity.common { color: #aaa; }
.card-rarity.uncommon { color: #2e7d32; }
.card-rarity.rare { color: #1565c0; }
.card-rarity.epic { color: #6a1b9a; }
.card-rarity.legendary { color: #f57f17; }

.card-stats {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.5rem;
}

.card-hashid {
  font-size: 0.7rem;
  color: #666;
  word-break: break-all;
  margin-bottom: 0.5rem;
}

.close-btn {
  display: block;
  margin: 0 auto;
  padding: 0.4rem 0.8rem;
  background: #42b983;
  color: #fff;
  border-radius: 4px;
  border: none;
  cursor: pointer;
}

/* 右侧卡牌描述 */
.card-profile {
  width: 250px;
  background: #fff;
  border-radius: 8px;
  padding: 0.5rem 1rem;
  max-height: 400px;      
  display: flex;
  flex-direction: column;
}

.profile-text {
  flex: 1;
  overflow-y: auto;       
  overflow-x: hidden;     
  white-space: normal;    
  word-break: break-word; 
  margin-top: 0.5rem;
  line-height: 1.4;
}
</style>
