<template>
  <div class="movie-card" :style="{ backgroundImage: `url(${bg})` }">
    <div class="movie-card__overlay"></div>

    <div class="movie-card__content">
      <div class="movie-card__header">
        <h3 class="movie-card__title">{{ title }}</h3>
        <button class="btn btn-outline movie-card__button" @click="onStart">
          开始挖矿
        </button>
      </div>

      <!-- 卡槽 -->
      <div class="slot-row">
        <div
          class="mining-slot"
          v-for="(slot, index) in slots"
          :key="index"
          @dragover.prevent
          @drop="onDrop(index)"
        >
          <template v-if="slot">
            <img :src="slot.avatar" class="slot-card" />
          </template>
          <template v-else>
            <span>拖拽卡牌到此</span>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import { request } from "../api/request";

export default {
  name: "DungeonRowCard",
  props: {
    title: String,
    bg: String,
    cardsPool: Array, // 父组件的卡牌池
    difficulty: String, 
  },
  inject: ['showAlert'], // <- 新增
  data() {
    return {
      slots: [null, null, null], // 每行3个槽位
    };
  },
 methods: {
  onDrop(index) {
    const card = this.$root.draggedCard;
    if (!card) return;

    // Vue 3 直接赋值即可，保证响应式
    this.slots[index] = card;

    // 清空全局拖拽状态
    this.$root.draggedCard = null;
  },
  async onStart() {
    if (!this.slots.every(slot => slot)) {
      this.showAlert("warning", "请先填满三个卡槽");
      return;
    }
      console.log('Injected showAlert:', this.showAlert);

    const payload = {
      difficulty: this.difficulty, 
      a: this.slots[0].hashid,
      b: this.slots[1].hashid,
      c: this.slots[2].hashid,

    };

    try {
      const res = await request("/auth/user/mine", {
        method: "POST",
        body: JSON.stringify(payload),
      });
      if (res.code === 200) {
      this.showAlert("success", "挑战成功！");
        this.slots = [null, null, null];
      } else {
        this.showAlert("danger", res.msg || "挖矿失败");
      }
    } catch (err) {
      console.error(err);
      this.showAlert("danger", "挖矿请求失败");
    }
  },
},

};
</script>

<style scoped>
/* 保留你的 movie-card 样式 */
@import url("https://fonts.googleapis.com/css?family=Montserrat:300,400,700,800");

.movie-card {
  background-size: contain;
  background-repeat: no-repeat;
  width: 100%;
  max-width: 800px;
  min-height: 180px;
  display: block;
  margin: 0.7rem auto;
  border-radius: 8px;
  box-shadow: 0px 8px 12px 0px rgba(0, 0, 0, 0.25);
  position: relative;

  
}
.movie-card__overlay {
  width: 100%;
  height: 100%;
  border-radius: 8px;
  background: linear-gradient(
    to right,
    rgba(42, 159, 255, 0.2) 0%,
    rgba(33, 33, 32, 1) 60%,
    rgba(33, 33, 32, 1) 100%
  );
  position: absolute;
  top: 0;
  left: 0;
}
.movie-card__content {
  width: 100%;
  max-width: 370px;
  display: flex;
  flex-direction: column;
  position: relative;
  float: right;
  padding-right: 1.2em;
  padding-bottom: 1em;
  z-index: 1;
}
.movie-card__header {
  display: flex;            /* ✅ 横向排列 */
  justify-content: space-between; 
  align-items: flex-start;  /* 顶部对齐 */
  margin-bottom: 0.5rem; 
}
.movie-card__title {
  color: #ffffff;
  margin-bottom: 0.25em;
  opacity: 0.85;
}
.movie-card__info {
  text-transform: uppercase;
  letter-spacing: 2px;
  font-size: 0.8em;
  color: #2a9fff;
  margin: 0;
  font-weight: 700;
  opacity: 0.7;
}
.movie-card__button {
  margin-top: 1rem;
}

.slot-row {
  display: flex;
  gap: 0.5rem;
}
.mining-slot {
  flex: 1;
  height: 100px;
  background: #e8f5e9;
  border: 2px dashed #aadadd;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  color: #333;
}
.slot-card {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 8px;
}

</style>
