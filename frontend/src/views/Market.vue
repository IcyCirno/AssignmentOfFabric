<template>
  <div class="market-page">
    <div class="walletBalanceCard">
  <div class="svgwrapper">
    <svg viewBox="0 0 24 26" fill="none" xmlns="http://www.w3.org/2000/svg">
      <rect
        x="0.539915"
        y="6.28937"
        width="21"
        height="4"
        rx="1.5"
        transform="rotate(-4.77865 0.539915 6.28937)"
        fill="#7D6B9D"
        stroke="black"
      ></rect>
      <circle
        cx="11.5"
        cy="5.5"
        r="4.5"
        fill="#E7E037"
        stroke="#F9FD50"
        stroke-width="2"
      ></circle>
      <path
        d="M2.12011 6.64507C7.75028 6.98651 12.7643 6.94947 21.935 6.58499C22.789 6.55105 23.5 7.23329 23.5 8.08585V24C23.5 24.8284 22.8284 25.5 22 25.5H2C1.17157 25.5 0.5 24.8284 0.5 24V8.15475C0.5 7.2846 1.24157 6.59179 2.12011 6.64507Z"
        fill="#BF8AEB"
        stroke="black"
      ></path>
      <path
        d="M16 13.5H23.5V18.5H16C14.6193 18.5 13.5 17.3807 13.5 16C13.5 14.6193 14.6193 13.5 16 13.5Z"
        fill="#BF8AEB"
        stroke="black"
      ></path>
    </svg>
  </div>

  <div class="balancewrapper">
    <span class="balanceHeading">Wallet balance</span>
    <p class="balance"><span id="currency">Gocoin：</span>{{balance}}</p>
  </div>

</div>
    <div class="page-grid">
      <!-- 左侧：市场卡牌 -->
      <div class="market-section">
        <h2>卡牌市场</h2>
        <div v-if="loading">加载中...</div>
        <div v-else-if="error" class="error">{{ error }}</div>
        <div v-else>
          <div class="card-grid">
            <div
              class="card-preview"
              v-for="card in displayedCards"
              :key="card.hashid"
              @click="openCard(card)"
            >
              <img :src="card.avatar" :alt="card.name" />
              <div class="card-name">{{ card.name || "未知卡牌" }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- 右侧：铸造区域 -->
      <div class="mint-section">
        <h2>铸造卡牌</h2>

        <form @submit.prevent="handleMint" class="mint-form">
          <label>
            卡牌名字：
            <input v-model="form.name" type="text" required />
          </label>

          <label>
            <BalanceSlider
            v-model="form.invest"
            :min="3"
            :max="balance"
            :step="1"
            :left-content="`投资: ${form.invest}`"
            :right-content="`余额: ${balance - form.invest}`"
          />
          </label>

          <button type="submit" class="learn-more">铸造</button>
        </form>

        <div v-if="showScratch || showCardModal" class="overlay">
          <ScratchCard
            v-if="showScratch"
            :width="450"
            :height="660"
            @complete="() => { showScratch = false; showCardModal = true }"
          >
            <img :src="newCard.avatar" alt="new card" class="scratch-card-img" />
          </ScratchCard>

          <CardModal_empty
            v-model="showCardModal"
            :card="newCard"
            :showActions="true"
            @close="closeCardModal"
          />
        </div>
      </div>
    </div>

    <!-- 市场卡牌弹窗 -->
    <CardModal
      v-model="showCardModalMarket"
      :card="selectedCard"
      :showActions="true"
      @close="closeCardModalMarket"
    />
  </div>
</template>

<script>
import { request } from "../api/request";
import { getProfile } from "../api/auth";
import CardModal from "../components/CardModal_sale.vue";
import CardModal_empty from "../components/CardModal_empty.vue";

import ScratchCard from "../components/ScratchToReveal.vue";
import BalanceSlider from "../components/BalanceSlider.vue"; // 假设你的滑动条组件路径

export default {
  components: { CardModal, ScratchCard, BalanceSlider, CardModal_empty },
  data() {
    return {
      // 市场数据
      cards: [],
      displayedCards: [],
      loading: true,
      error: "",
      showCardModalMarket: false,
      selectedCard: null,

      // 铸造数据
      balance: 0,
      form: { name: "", invest: "", avatar: "" },
      newCard: {},
      showCardModal: false,
      showScratch: false
    };
  },
  async mounted() {
    await this.fetchMarketCards();
    await this.fetchBalance();
  },
  methods: {
    /** 市场 */
    async fetchMarketCards() {
      try {
        const res = await request("/auth/market/query", { method: "POST" });
        if (res.code === 200) {
          const cardsMap = {};
          (res.data || []).forEach(t => {
            if (t.onsale) {
              cardsMap[t.hashid] = {
                ...t,
                price: t.price || t.cost || 0,
                seller: t.seller || t.owner,
                transID: t.trans_id || t.transID || ""
              };
            }
          });
          this.displayedCards = Object.values(cardsMap);
        } else {
          this.error = res.msg || "加载失败";
        }
      } catch (e) {
        console.error("拉取市场卡牌失败:", e);
        this.error = "服务器错误";
      } finally {
        this.loading = false;
      }
    },
    openCard(card) {
      this.selectedCard = card;
      this.showCardModalMarket = true;
    },
    closeCardModalMarket() {
      this.showCardModalMarket = false;
      this.selectedCard = null;
    },

    /** 铸造 */
    async fetchBalance() {
      try {
        const res = await getProfile();
        if (res.code === 200) this.balance = res.data.gocoin;
      } catch (e) {
        console.error(e);
      }
    },
    async handleMint() {
      if (!this.form.name || !this.form.invest) return;
      const payload = {
        name: this.form.name,
        invest: parseInt(this.form.invest, 10)
      };
      try {
        const res = await request("/auth/card/mint", {
          method: "POST",
          body: JSON.stringify(payload)
        });
        if (res.code === 200) {
          this.newCard = res.data;
          this.showScratch = true;
          this.form = { name: "", invest: "", avatar: "" };
        }
      } catch (e) {
        console.error(e);
      }
    },
    closeCardModal() {
      this.showCardModal = false;
      this.newCard = {};
    }
  }
};
</script>

<style scoped>
.page-grid {
  display: flex;
  gap: 3rem;
  height: calc(100vh - 8rem); /* 让左右区域撑满页面高度 */
  padding:6rem 2rem; /* 页面整体上下/左右间距 */
  
}

/* 左侧市场滚动 */
.market-section {
  flex: 2;
  overflow-y: auto;   /* 可以垂直滚动 */
  padding-right: 0.5rem; /* 防止滚动条覆盖内容 */
  padding: 1rem;       /* 内边距 */
  border-radius: 8px;  /* 圆角 */
  background: #f9fafb; /* 背景颜色和右侧一致 */
}

/* 右侧铸造固定 */
.mint-section {
  flex: 1;
  background: #f9fafb;
  padding: 1rem;
  border-radius: 8px;
  height: 40%; /* 固定高度撑满容器 */
  position: sticky; /* 页面滚动时保持可视 */
  top: 1rem;        /* 离顶部留一点间距 */
}

/* 保留原市场样式 */
.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 1rem;
}
.card-preview {
  position: relative;
  width: 180px;
  height: 240px;
  overflow: hidden;
  border-radius: 8px;
  cursor: pointer;
  display: flex;
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
.card-preview:hover { transform: scale(1.05); }
.card-name {
  width: 100%;
  text-align: center;
  background: rgba(0,0,0,0.4);
  color: white;
  font-weight: bold;
  padding: 0.2rem 0;
  position: absolute;
  bottom: 0;
  left: 0;
  border-radius: 0 0 8px 8px;
}

/* 铸造表单样式 */
.mint-form label { display: block; margin-bottom: 1rem; }
.mint-form input, .mint-form textarea { width: 100%; padding: 0.5rem; margin-top: 0.3rem; }
/* From Uiverse.io by Madflows */ 
button {
 position: relative;
 display: inline-block;
 cursor: pointer;
 outline: none;
 border: 0;
 vertical-align: middle;
 text-decoration: none;
 font-family: inherit;
 font-size: 15px;
}

button.learn-more {
 font-weight: 600;
 color: #382b22;
 text-transform: uppercase;
 padding: 1.25em 2em;
 background: #87cce2;
 border: 2px solid #b18597;
 border-radius: 0.75em;
 -webkit-transform-style: preserve-3d;
 transform-style: preserve-3d;
 -webkit-transition: background 150ms cubic-bezier(0, 0, 0.58, 1), -webkit-transform 150ms cubic-bezier(0, 0, 0.58, 1);
 transition: transform 150ms cubic-bezier(0, 0, 0.58, 1), background 150ms cubic-bezier(0, 0, 0.58, 1), -webkit-transform 150ms cubic-bezier(0, 0, 0.58, 1);
 display: block; 
 margin: 1rem auto 0 auto; /* 上间距 1rem，左右居中，下间距 0 */
}

button.learn-more::before {
 position: absolute;
 content: '';
 width: 100%;
 height: 100%;
 top: 0;
 left: 0;
 right: 0;
 bottom: 0;
 background: #0e61b5;
 border-radius: inherit;
 -webkit-box-shadow: 0 0 0 2px #061251, 0 0.625em 0 0 #ffe3e2;
 box-shadow: 0 0 0 2px #8596b1, 0 0.625em 0 0 hsl(258, 59%, 93%);
 -webkit-transform: translate3d(0, 0.75em, -1em);
 transform: translate3d(0, 0.75em, -1em);
 transition: transform 150ms cubic-bezier(0, 0, 0.58, 1), box-shadow 150ms cubic-bezier(0, 0, 0.58, 1), -webkit-transform 150ms cubic-bezier(0, 0, 0.58, 1), -webkit-box-shadow 150ms cubic-bezier(0, 0, 0.58, 1);
}

button.learn-more:hover {
 background: #ffe9e9;
 -webkit-transform: translate(0, 0.25em);
 transform: translate(0, 0.25em);
}

button.learn-more:hover::before {
 -webkit-box-shadow: 0 0 0 2px #b18597, 0 0.5em 0 0 #ffe3e2;
 box-shadow: 0 0 0 2px #b18597, 0 0.5em 0 0 #ffe3e2;
 -webkit-transform: translate3d(0, 0.5em, -1em);
 transform: translate3d(0, 0.5em, -1em);
}

button.learn-more:active {
 background: #ffe9e9;
 -webkit-transform: translate(0em, 0.75em);
 transform: translate(0em, 0.75em);
}

button.learn-more:active::before {
 -webkit-box-shadow: 0 0 0 2px #b18597, 0 0 #ffe3e2;
 box-shadow: 0 0 0 2px #b18597, 0 0 #ffe3e2;
 -webkit-transform: translate3d(0, 0, -1em);
 transform: translate3d(0, 0, -1em);
}

.overlay {
  position: fixed;
  top:0; left:0;
  width:100%; height:100%;
  background: rgba(0,0,0,0.5);
  display:flex; justify-content:center; align-items:center;
  z-index:2000;
}
.scratch-card-img {
  width: 100%; height: 100%; object-fit: cover; border-radius: 5px;
}

/* From Uiverse.io by vinodjangid07 */ 
.walletBalanceCard {
  position: fixed;
  top: 4rem;         /* 上边距更大 */
  right: 1rem;       /* 靠右 */
  z-index: 3000;     /* 保证不被挡住 */

  width: fit-content;
  height: 55px;
  background-color: hsl(211, 39%, 51%);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 12px;
  padding: 0px 12px;
  font-family: Arial, Helvetica, sans-serif;
}
.svgwrapper {
  width: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.svgwrapper svg {
  width: 100%;
}
.balancewrapper {
  display: flex;
  align-items: flex-start;
  justify-content: flex-start;
  flex-direction: column;
  width: 120px;
  gap: 0px;
}
.balanceHeading {
  font-size: 8px;
  color: rgb(214, 214, 214);
  font-weight: 100;
  letter-spacing: 0.6px;
}
.balance {
  font-size: 13.5px;
  color: white;
  font-weight: 600;
  letter-spacing: 0.5px;
}

.plussign {
  font-size: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
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
