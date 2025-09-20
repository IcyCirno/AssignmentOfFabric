<template>
  <div class="dungeon-page">
    <!-- 左侧：卡牌收藏区 -->
    <aside class="cards-collection">
      <h3>空闲卡牌</h3>
      <div class="card-grid">
        <div
          class="card-preview"
          v-for="card in myCards"
          :key="card.hashid"
          
          @click="openCard(card)"
          @dragstart="onDragStart(card)"
        >
          <img :src="card.avatar" :alt="card.name" />
          <div class="card-name">{{ card.name || '未知卡牌' }}</div>
        </div>
      </div>
    </aside>
    <CardModal
      v-model="showCardModal"
      :card="selectedCard"
      :showActions="true"
      
      @close="closeCardModal"
    />
    <!-- 右侧：挖矿副本区 -->
    <main class="dungeon-panel">
      <DungeonRowCard
        v-for="(bg, index) in bgList"
        :key="index"
        :title="`副本难度 ${index + 1}`"
        :bg="bg"
        :cards-pool="myCards"
        :difficulty="index === 0 ? 'simple' : index === 1 ? 'common' : 'hard'"
      />
    </main>
  </div>
</template>


<script>
import { request } from "../api/request";
import DungeonRowCard from "../components/DungeonRowCard.vue";
import CardModal from "../components/CardModal_empty.vue";

import PicEasy from "../assets/Pic_easy.jpg";
import PicNorm from "../assets/Pic_norm.jpg";
import PicHard from "../assets/Pic_hard.jpg";

export default {
  name: "MinePage",
  components: { DungeonRowCard, CardModal },
  data() {
    return {
      myCards: [], // 卡牌池
      bgList: [PicEasy, PicNorm, PicHard],
      showCardModal: false,
      selectedCard: null,
    };
  },
  async mounted() {
    await this.fetchCards();
  },
  methods: {
    async fetchCards() {
      try {
        const res = await request("/auth/card/query", { method: "POST" });
        if (res.code === 200) {
          // res.data = res.data.filter((c) => c.owner === this.user.name);
          this.myCards = res.data.filter((c) => !c.destroy && !c.onsale);
        }
      } catch (err) {
        console.error("获取卡牌失败", err);
      }
    },
    onDragStart(card) {
      // 可以通过全局事件或 provide/inject 传给子组件
      this.$root.draggedCard = card;
    },
    openCard(card) {
      this.selectedCard = card;
      this.showCardModal = true;
    },
  },
};
</script>

<style scoped>
.dungeon-page {
  display: flex;
  gap: 1rem;
  padding: 1rem;
  height: 85vh;
  box-sizing: border-box;
}

/* 左侧卡牌区域 */
.cards-collection {
  width: 40%;
  background: #f0f0f0;
  border-radius: 8px;
  padding: 1rem;
  overflow-y: auto;
  margin-top: 2rem; /* ✅ 增加顶部距离 */

}
.cards-collection h3 {
  margin-bottom: 1rem;
}
.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(100px, 1fr));
  gap: 0.5rem;
}
.card-preview {
  position: relative;
  width: 100%;
  height: 140px;
  overflow: hidden;
  border-radius: 8px;
  cursor: grab;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  align-items: center;
  background: #fff;
  border: 1px solid #ccc;
}
.card-preview img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.card-name {
  width: 100%;
  text-align: center;
  background: rgba(0, 0, 0, 0.4);
  color: white;
  font-weight: bold;
  padding: 0.2rem 0;
  position: absolute;
  bottom: 0;
  left: 0;
  border-radius: 0 0 8px 8px;
}

/* 右侧挖矿副本区域 */
.dungeon-panel {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4.5rem;
  overflow-y: auto;
  margin-top: 2.3rem;

}
.dungeon-row {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}
.mining-slot {
  flex: 1;
  height: 120px;
  background: #e8f5e9;
  border: 2px dashed #4caf50;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  text-align: center;
}
.slot-card {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 8px;
}
.start-btn {
  padding: 0.5rem 1rem;
  background: #4caf50;
  color: #fff;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: bold;
  transition: transform 0.2s;
}
.start-btn:hover {
  transform: scale(1.05);
}
</style>
