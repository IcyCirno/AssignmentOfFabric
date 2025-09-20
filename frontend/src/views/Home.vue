<template>
  <div class="home-page">
    
    <aside class="profile-panel">
      <div class="glitch-profile-wrapper">
    <div class="glitch-card profile-card">
      <div class="card-header-visual"></div>

      <!-- 用户头像 -->
      <div class="profile-avatar">
        <svg
          class="octocat-svg"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 16 16"
        >
          <path
            fill-rule="evenodd"
            d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 
               5.47 7.59.4.07.55-.17.55-.38 
               0-.19-.01-.82-.01-1.49
               -2.01.37-2.53-.49-2.69-.94
               -.09-.23-.48-.94-.82-1.13
               -.28-.15-.68-.52-.01-.53
               .63-.01 1.08.58 1.23.82
               .72 1.21 1.87.87 2.33.66
               .07-.52.28-.87.51-1.07
               -1.78-.2-3.64-.89-3.64-3.95
               0-.87.31-1.59.82-2.15
               -.08-.2-.36-1.02.08-2.12
               0 0 .67-.21 2.2.82
               .64-.18 1.32-.27 2-.27
               .68 0 1.36.09 2 .27
               1.53-1.04 2.2-.82 2.2-.82
               .44 1.1.16 1.92.08 2.12
               .51.56.82 1.27.82 2.15
               0 3.07-1.87 3.75-3.65 3.95
               .29.25.54.73.54 1.48
               0 1.07-.01 1.93-.01 2.2
               0 .21.15.46.55.38
               A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"
          ></path>
        </svg>
      </div>

      <div class="card-body profile-body">
        <!-- 用户信息 -->
        <div class="profile-info">
          <div class="profile-name" :data-text="user.name">
            {{ user.name || '未登录' }}
          </div>
          <p class="profile-title">GoCoin: {{ user.gocoin || 0 }}</p>
        </div>

        <!-- 统计信息：卡牌数量 -->
        <div class="profile-stats">
          <div class="stat-item">
            <span class="stat-label">我的卡牌</span>
            <span class="stat-value" :data-text="myCards.length">
              {{ myCards.length }}
            </span>
          </div>
          <div class="stat-item">
            <span class="stat-label">上架中</span>
            <span class="stat-value" :data-text="saleCards.length">
              {{ saleCards.length }}
            </span>
          </div>
        </div>

        <!-- 按钮 -->
        <button class="submit-btn" @click="handleLogout" data-text="退出登录">
          <span class="btn-text">退出登录</span>
        </button>
      </div>
    </div>
  </div>
    <div class="bottom-panel">

        <!-- From Uiverse.io by Juanes200122 --> 
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
            stroke=""
            fill="#414750"
            points="70,70 150,50 130,130 50,150"
          ></polygon>
          <polygon
            stroke-width="2"
            stroke=""
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
            stroke=""
            fill="#b7870f"
            points="80,50 80,75 80,99 40,75"
          ></polygon>
          <polygon
            transform="translate(20, 31)"
            stroke-width="2"
            stroke=""
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
            stroke=""
            fill="#d3a410"
            points="80,50 80,75 80,99 40,75"
          ></polygon>
          <polygon
            transform="rotate(0 100 100) translate(60, 20)"
            stroke-width="2"
            stroke=""
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
            stroke=""
            fill="#ffe4a1"
            points="5,0 5,5 0,5 0,0"
            id="particles"
          ></polygon>
          <polygon
            transform="rotate(45 100 100) translate(80, 55)"
            stroke-width="2"
            stroke=""
            fill="#ccb069"
            points="6,0 6,6 0,6 0,0"
            id="particles"
          ></polygon>
          <polygon
            transform="rotate(45 100 100) translate(70, 80)"
            stroke-width="2"
            stroke=""
            fill="#fff"
            points="2,0 2,2 0,2 0,0"
            id="particles"
          ></polygon>
          <polygon
            stroke-width="2"
            stroke=""
            fill="#292d34"
            points="29.5,99.8 100,142 100,172 29.5,130"
          ></polygon>
          <polygon
            transform="translate(50, 92)"
            stroke-width="2"
            stroke=""
            fill="#1f2127"
            points="50,50 120.5,8 120.5,35 50,80"
          ></polygon>
        </g>
      </svg>


      </div>
    </aside>

    <!-- 右侧 cards 区域 -->
    <main class="cards-panel">
      <!-- 收藏区 -->
      <section class="collection-section">
        <h3>我的收藏</h3>
        <div class="card-grid">
          <div
            class="card-preview"
            v-for="card in myCards"
            :key="card.hashid"
            draggable="true"
            @click="openCard(card)"
          >
            <img :src="card.avatar" :alt="card.name" />
            <div class="card-name">{{ card.name || '未知卡牌' }}</div>
          </div>
        </div>
      </section>

      <!-- 售卖区 -->
      <section class="sale-section">
        <h3>我的上架</h3>
        <div class="card-grid">
          <div
            class="card-preview"
            v-for="card in saleCards"
            :key="card.hashid"
            @click="openCard(card)"
          >
            <img :src="card.avatar" :alt="card.name" />
            <div class="card-name">{{ card.name }}</div>
          </div>
        </div>
      </section>
    </main>

    <!-- 卡牌详情弹窗 -->
    <CardModal
      v-model="showCardModal"
      :card="selectedCard"
      :showActions="true"
      
      @close="closeCardModal"
    />
  </div>
</template>

<script>
import { getProfile, logout } from "../api/auth";
import { request } from "../api/request";
import CardModal from "../components/CardModal.vue";

export default {
  components: { CardModal },
  data() {
    return {
      user: {},
      loading: true,
      error: "",
      myCards: [],
      saleCards: [],
      showCardModal: false,
      selectedCard: null,
      // 挖矿卡槽
    };
  },
  async mounted() {
    await this.fetchProfile();
    await this.fetchCards();
  },
  methods: {
    async fetchProfile() {
      try {
        const token = localStorage.getItem("token");
        if (!token) return (window.location.href = "/auth");
        const res = await getProfile();
        if (res.code === 200) this.user = res.data;
        else this.error = res.msg || "加载失败";
      } catch (err) {
        this.error = "服务器请求失败";
        console.error(err);
      } finally {
        this.loading = false;
      }
    },
    handleLogout() {
      logout(this.$router);
    },
    async fetchCards() {
      try {
        const res = await request("/auth/card/query", { method: "POST" });
        if (res.code === 200) {
           let filtered = res.data.filter(c => !c.destroy && c.owner === this.user.name);
            const seen = new Set();
            filtered = filtered.filter(c => {
              if (seen.has(c.hashid)) return false;
              seen.add(c.hashid);
              return true;
            });

            // 分别赋值收藏区和上架区
            this.myCards = filtered;
            this.saleCards = filtered.filter(c => c.onsale);
        } else this.$root.showAlert("danger", "获取卡牌失败");
      } catch (e) {
        console.error(e);
      }
    },
    openCard(card) {
      this.selectedCard = card;
      this.showCardModal = true;
    },
    closeCardModal() {
      this.showCardModal = false;
      this.selectedCard = null;
    },
    
  
  }
};
</script>

<style scoped>

/* From Uiverse.io by Juanes200122 */ 
.container {
  background-color: #414141;
}
@keyframes bounce {
  0%,
  100% {
    translate: 0px 36px;
  }
  50% {
    translate: 0px 46px;
  }
}
@keyframes bounce2 {
  0%,
  100% {
    translate: 0px 46px;
  }
  50% {
    translate: 0px 56px;
  }
}

@keyframes umbral {
  0% {
    stop-color: #d3a5102e;
  }
  50% {
    stop-color: rgba(211, 165, 16, 0.519);
  }
  100% {
    stop-color: #d3a5102e;
  }
}
@keyframes partciles {
  0%,
  100% {
    translate: 0px 16px;
  }
  50% {
    translate: 0px 6px;
  }
}
#particles {
  animation: partciles 4s ease-in-out infinite;
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


.home-page {
  display: flex;
  height: 85vh;
  gap: 1rem;
  padding: clamp(1rem, 2vw, 6rem);
  align-self: stretch
}
.profile-panel {
  width: 220px;
  background: #fff;
  border-radius: 8px;
  padding: 1rem;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  /* height: 100%; */
}
.mining-slot {
  height: 180px;
  margin-top: 1rem;
  background: #f5f5f5;
  border: 2px dashed #ccc;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  text-align: center;
  transition: background 0.3s;
}
.mining-slot.has-card {
  border: 2px solid #4caf50;
  background: #e8f5e9;
}
.bottom-panel {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  align-items: center; /* 水平居中 */
}


.slot-card {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 8px;
}
.clear-btn {
  position: absolute;
  top: 4px;
  right: 4px;
  padding: 2px 6px;
  font-size: 12px;
  cursor: pointer;
}
.cards-panel {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 1rem;

}
.collection-section {
  flex: 2.2;
  background: #f0f0f0;
  border-radius: 8px;
  padding: 1rem;
  display: flex;
     border-top: 2px solid #ccc; 

  flex-direction: column;
}
.sale-section  {
  flex: 1;
  background: #f0f0f0;
  border-radius: 8px;
  padding: 1rem;
  display: flex;
   border-top: 2px solid #ccc; 
  flex-direction: column;

}
.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  gap: 1rem;
  flex: 1;
}
.card-preview {
  position: relative;
  width: 150px;
  height: 200px;
  overflow: hidden;
  border-radius: 8px;
  cursor: pointer;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  align-items: center;
  background: #f5f5f5;
  transition: transform 0.2s;
}
.card-preview img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.card-preview:hover {
  transform: scale(1.05);
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

/* From Uiverse.io by pharmacist-sabot */ 
/* --- Root Variables & Wrapper --- */
.glitch-profile-wrapper {
  --bg-color: #61b7bf;
  --primary-color: #000000;
  --secondary-color: #9865c7;
  --text-color: #000000;
  --font-family: "Fira Code", Consolas, "Courier New", Courier, monospace;
  --glitch-anim-duration: 0.4s;

  display: flex;
  justify-content: center;
  align-items: center;
  font-family: var(--font-family);
  padding: 1rem;
}

/* --- Card Structure --- */
.glitch-card {
  position: relative;
  background-color: var(--bg-color);
  width: 19rem;
  border-radius: 1rem;
  border: 1px solid rgba(0, 242, 234, 0.2);
  box-shadow:
    0 0 1.5rem rgba(0, 242, 234, 0.1),
    inset 0 0 1rem rgba(0, 0, 0, 0.5);
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.card-header-visual {
  height: 3.5rem;
  background: linear-gradient(
      45deg,
      rgba(0, 242, 234, 0.1),
      rgba(13, 13, 13, 0.8)
    ),
    radial-gradient(circle, rgba(168, 85, 247, 0.05) 0%, transparent 70%);
  border-bottom: 1px solid rgba(0, 242, 234, 0.1);
}

/* --- Avatar --- */
.profile-avatar {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  top: 0.75rem;
  z-index: 2;
  width: 3.5rem;
  height: 3.5rem;
  background: var(--bg-color);
  border-radius: 50%;
  padding: 0.9rem;
  box-shadow: 0 0 1.5rem rgba(0, 0, 0, 0.5);
  border: 1px solid rgba(0, 242, 234, 0.3);
  transition: background-color 0.5s ease-in-out;
}

.octocat-svg {
  width: 100%;
  height: 100%;
  fill: var(--primary-color);
  transition: all 0.6s cubic-bezier(0.4, 0, 0.2, 1);
  filter: drop-shadow(0 0 10px rgba(0, 242, 234, 0.3));
}

.profile-avatar:hover {
  background-color: var(--primary-color);
}
.profile-avatar:hover .octocat-svg {
  fill: var(--bg-color);
  transform: scale(1.1) rotate(-5deg);
}

/* --- Card Body --- */
.card-body {
  padding: 1.1rem;
  padding-top: 0;
  background-color: var(--bg-color);
  display: flex;
  flex-direction: column;
  flex-grow: 1;
}

.profile-body {
  align-items: center;
  text-align: center;
  padding-top: 2.75rem;
  display: flex;
  flex-direction: column;
  flex-grow: 1;
}

.profile-info {
  margin-bottom: 1.1rem;
}
.profile-name {
  font-size: 1.6rem;
  font-weight: 700;
  color: var(--text-color);
  text-transform: uppercase;
  letter-spacing: 0.1em;
  margin: 0;
  position: relative;
  cursor: pointer;
}
.profile-title {
  font-size: 0.75rem;
  color: var(--primary-color);
  opacity: 0.8;
  margin: 0.25rem 0 0 0;
  letter-spacing: 0.1em;
}

/* --- GitHub Stats --- */
.profile-stats {
  width: 100%;
  margin-top: auto;
  padding-top: 1rem;
  border-top: 1px solid rgba(0, 242, 234, 0.2);
  display: flex;
  justify-content: space-around;
}
.stat-item {
  text-align: center;
  cursor: pointer;
}
.stat-label {
  font-size: 0.65rem;
  font-weight: 700;
  color: var(--text-color);
  opacity: 0.6;
  letter-spacing: 0.1em;
}
.stat-value {
  display: block;
  font-size: 1.4rem;
  font-weight: 700;
  color: var(--primary-color);
  margin-top: 0.25rem;
  position: relative;
}

/* --- Button --- */
.submit-btn {
  display: block;
  width: 100%;
  margin-top: 1.1rem;
  padding: 0.7em 1em;
  background-color: transparent;
  border: 2px solid var(--primary-color);
  color: var(--primary-color);
  font-family: inherit;
  font-size: 0.9rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.2em;
  cursor: pointer;
  position: relative;
  transition: all 0.3s;
  overflow: hidden;
  text-align: center;
  text-decoration: none;
}

.submit-btn:hover,
.submit-btn:focus {
  background-color: var(--primary-color);
  color: var(--bg-color);
  box-shadow: 0 0 1.5rem var(--primary-color);
  outline: none;
}

.submit-btn:active {
  transform: scale(0.97);
}
.submit-btn .btn-text {
  position: relative;
  z-index: 1;
  transition: opacity 0.2s ease;
}

/* --- Glitch Effects --- */
.profile-name:hover::before,
.profile-name:hover::after,
.stat-value:hover::before,
.stat-value:hover::after,
.submit-btn:hover::before,
.submit-btn:hover::after,
.submit-btn:focus::before,
.submit-btn:focus::after {
  content: attr(data-text);
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}
.profile-name:hover,
.stat-item:hover .stat-value {
  color: transparent;
}
.submit-btn:hover .btn-text {
  opacity: 0;
}
.profile-name:hover::before,
.stat-value:hover::before,
.submit-btn:hover::before,
.submit-btn:focus::before {
  opacity: 1;
  color: var(--secondary-color);
  animation: glitch-anim var(--glitch-anim-duration)
    cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
}
.profile-name:hover::after,
.stat-value:hover::after {
  background-color: var(--bg-color);
  color: var(--primary-color);
  animation: glitch-anim var(--glitch-anim-duration)
    cubic-bezier(0.25, 0.46, 0.45, 0.94) reverse both;
}

.submit-btn:hover::after,
.submit-btn:focus::after {
  color: var(--bg-color);
  animation: glitch-anim var(--glitch-anim-duration)
    cubic-bezier(0.25, 0.46, 0.45, 0.94) reverse both;
}

@keyframes glitch-anim {
  0% {
    transform: translate(0);
    clip-path: inset(0 0 0 0);
  }
  20% {
    transform: translate(-0.3rem, 0.15rem);
    clip-path: inset(50% 0 20% 0);
  }
  40% {
    transform: translate(0.15rem, -0.1rem);
    clip-path: inset(20% 0 60% 0);
  }
  60% {
    transform: translate(-0.25rem, 0.1rem);
    clip-path: inset(80% 0 5% 0);
  }
  80% {
    transform: translate(0.25rem, -0.15rem);
    clip-path: inset(30% 0 45% 0);
  }
  100% {
    transform: translate(0);
    clip-path: inset(0 0 0 0);
  }
}


body, html {
  margin: 0;
  height: 100%;
  width: 100%;
  background: #000; /* 外部黑边 */
  display: flex;
  justify-content: center;
  align-items: center;
}

.app-wrapper {
  aspect-ratio: 16 / 9;
  width: 100%;
  max-height: 100%;
  background: #fff; /* 你的网页内容区域 */
  overflow: hidden;
}

</style>
