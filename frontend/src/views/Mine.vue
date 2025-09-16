<template>
  <div class="flex min-h-screen bg-gray-100 p-8 gap-8">
    <!-- 左边卡槽 -->
    <div class="w-1/2 flex items-center justify-center">
    <div class="card-slot">
        <!-- 冷却中 -->
        <img
        v-if="cooldown"
        src="../assets/cooldown.png"
        alt="冷却中"
        class="card-image"
        />
        <!-- 已选卡 -->
        <img
        v-else-if="selectedCard"
        :src="selectedCard.avatar"
        :alt="selectedCard.name"
        class="card-image"
        />
        <!-- 空卡槽 -->
        <div v-else class="card-empty" @click="goCollection">
        去选卡
        </div>
    </div>
    </div>


    <!-- 右边信息区 -->
    <div class="w-1/2 flex flex-col gap-6">
      <!-- 上半部分：SVG 动画 -->
      <div class="bg-white rounded-2xl shadow p-6 flex justify-center items-center">
        <svg xmlns="http://www.w3.org/2000/svg" height="200" width="200">
          <g style="order: -1;">
            <polygon
              transform="rotate(45 100 100)"
              stroke-width="1"
              stroke="#d3a410"
              fill="none"
              points="70,70 148,50 130,130 50,150"
              id="bounce"
            ></polygon>
            <polygon
              transform="rotate(45 100 100)"
              stroke-width="1"
              stroke="#d3a410"
              fill="none"
              points="70,70 148,50 130,130 50,150"
              id="bounce2"
            ></polygon>
            <polygon
              transform="rotate(45 100 100)"
              stroke-width="2"
              fill="#414750"
              points="70,70 150,50 130,130 50,150"
            ></polygon>
            <polygon
              stroke-width="2"
              fill="url(#gradiente)"
              points="100,70 150,100 100,130 50,100"
            ></polygon>
            <defs>
              <linearGradient y2="100%" x2="10%" y1="0%" x1="0%" id="gradiente">
                <stop style="stop-color: #1e2026;stop-opacity:1" offset="20%"></stop>
                <stop style="stop-color:#414750;stop-opacity:1" offset="60%"></stop>
              </linearGradient>
            </defs>
            <polygon
              transform="translate(20, 31)"
              stroke-width="2"
              fill="#b7870f"
              points="80,50 80,75 80,99 40,75"
            ></polygon>
            <polygon
              transform="translate(20, 31)"
              stroke-width="2"
              fill="url(#gradiente2)"
              points="40,-40 80,-40 80,99 40,75"
            ></polygon>
            <defs>
              <linearGradient y2="100%" x2="0%" y1="-17%" x1="10%" id="gradiente2">
                <stop style="stop-color: #d3a51000;stop-opacity:1" offset="20%"></stop>
                <stop
                  style="stop-color:#d3a51054;stop-opacity:1"
                  offset="100%"
                  id="animatedStop"
                ></stop>
              </linearGradient>
            </defs>
            <polygon
              transform="rotate(180 100 100) translate(20, 20)"
              stroke-width="2"
              fill="#d3a410"
              points="80,50 80,75 80,99 40,75"
            ></polygon>
            <polygon
              transform="rotate(0 100 100) translate(60, 20)"
              stroke-width="2"
              fill="url(#gradiente3)"
              points="40,-40 80,-40 80,85 40,110.2"
            ></polygon>
            <defs>
              <linearGradient y2="100%" x2="10%" y1="0%" x1="0%" id="gradiente3">
                <stop style="stop-color: #d3a51000;stop-opacity:1" offset="20%"></stop>
                <stop
                  style="stop-color:#d3a51054;stop-opacity:1"
                  offset="100%"
                  id="animatedStop"
                ></stop>
              </linearGradient>
            </defs>
            <polygon
              transform="rotate(45 100 100) translate(80, 95)"
              stroke-width="2"
              fill="#ffe4a1"
              points="5,0 5,5 0,5 0,0"
              id="particles"
            ></polygon>
            <polygon
              transform="rotate(45 100 100) translate(80, 55)"
              stroke-width="2"
              fill="#ccb069"
              points="6,0 6,6 0,6 0,0"
              id="particles"
            ></polygon>
            <polygon
              transform="rotate(45 100 100) translate(70, 80)"
              stroke-width="2"
              fill="#fff"
              points="2,0 2,2 0,2 0,0"
              id="particles"
            ></polygon>
            <polygon
              stroke-width="2"
              fill="#292d34"
              points="29.5,99.8 100,142 100,172 29.5,130"
            ></polygon>
            <polygon
              transform="translate(50, 92)"
              stroke-width="2"
              fill="#1f2127"
              points="50,50 120.5,8 120.5,35 50,80"
            ></polygon>
          </g>
        </svg>
      </div>

      <!-- 下半部分：挖矿信息 -->
      <div class="bg-white rounded-2xl shadow p-6 flex flex-col gap-4">
        <h2 class="text-xl font-bold">挖矿信息</h2>
        <p>当前余额: {{ gocoin }} GoCoin</p>
        <p>
          挖矿速度:
          <span v-if="selectedCard">{{ selectedCard.rarity * 5 }} GoCoin/min</span>
          <span v-else>未选择卡牌</span>
        </p>
        <p class="text-gray-500">
          {{ cooldown ? "冷却中，请稍后再试" : "准备就绪" }}
        </p>

        <button
          @click="handleMine"
          :disabled="cooldown || !selectedCard"
          :class="[
            'px-6 py-3 rounded-xl text-white font-bold mt-4',
            cooldown || !selectedCard
              ? 'bg-gray-400 cursor-not-allowed'
              : 'bg-green-500 hover:bg-green-600'
          ]"
        >
          {{ cooldown ? "冷却中" : "开始挖矿" }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const token = localStorage.getItem("token");

const selectedCard = ref(null); // {id, name, avatar, rarity}
const cooldown = ref(false);
const gocoin = ref(0);

async function fetchUser() {
  try {
    const res = await fetch("/api/auth/user/info", {
      headers: { Authorization: `Bearer ${token}` },
    });
    const data = await res.json();
    if (data.code === 200) {
      gocoin.value = data.data.gocoin;
      if (data.data.cooldown_end) {
        const endTime = new Date(data.data.cooldown_end).getTime();
        if (Date.now() < endTime) {
          cooldown.value = true;
        }
      }
    }
  } catch (err) {
    console.error("获取用户信息失败", err);
  }
}

async function handleMine() {
  if (!selectedCard.value) {
    this.$root.showAlert("warning", "请先选一张卡牌");
    return;
  }

  try {
    const res = await fetch("/api/auth/user/mine", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ hash_id: selectedCard.value.id }),
    });
    const data = await res.json();
    if (data.code === 200) {
        this.$root.showAlert("success", "挖矿成功！");
      cooldown.value = true;
      gocoin.value += 10; // 临时加，后端可返回收益
    } else {
        this.$root.showAlert("danger", "挖矿失败！");
    }
  } catch (err) {
    console.error("挖矿失败", err);
  }
}

function goCollection() {
  router.push("/collection");
}

onMounted(() => {
  fetchUser();
  const savedCard = localStorage.getItem("selectedCard");
  if (savedCard) {
    selectedCard.value = JSON.parse(savedCard);
  }
});
</script>

.card-slot {
  width: 190px;
  height: 254px;
  background: rgb(236, 236, 236);
  box-shadow: rgba(0, 0, 0, 0.4) 0px 2px 4px,
              rgba(0, 0, 0, 0.3) 0px 7px 13px -3px,
              rgba(0, 0, 0, 0.2) 0px -3px 0px inset;
  border-radius: 12px;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
}

/* 卡牌或冷却图片填充框架 */
.card-image {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

/* 空卡槽显示文字 */
.card-empty {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6b7280; /* gray-500 */
  font-size: 1rem;
  cursor: pointer;
  user-select: none;
}


<style scoped>
/* SVG 动画样式 */
@keyframes bounce {
  0%, 100% { translate: 0px 36px; }
  50% { translate: 0px 46px; }
}
@keyframes bounce2 {
  0%, 100% { translate: 0px 46px; }
  50% { translate: 0px 56px; }
}
@keyframes umbral {
  0% { stop-color: #d3a5102e; }
  50% { stop-color: rgba(211, 165, 16, 0.519); }
  100% { stop-color: #d3a5102e; }
}
@keyframes particles {
  0%, 100% { translate: 0px 16px; }
  50% { translate: 0px 6px; }
}
#particles {
  animation: particles 4s ease-in-out infinite;
}
#animatedStop {
  animation: umbral 4s infinite;
}
#bounce {
  animation: bounce 4s ease-in-out infinite;
  translate: 0px 36px;
}
#bounce2 {
  animation: bounce2 4s ease-in-out infinite;
  translate: 0px 46px;
  animation-delay: 0.5s;
}

</style>
