<template>
  <div class="collection-page">
    <h2>我的收藏</h2>

    <!-- 筛选与排序 -->
    <div class="controls">
      <label>
        排序：
        <select v-model="sortKey" @change="applyFilterAndSort">
          <option value="rarity">按稀有度</option>
          <option value="name">按名称</option>
        </select>
      </label>
      <label>
        <input type="checkbox" v-model="filteronsale" @change="applyFilterAndSort" />
        只显示上架卡牌
      </label>
    </div>

    <!-- 卡牌网格 -->
    <div class="card-grid">
      <div 
        class="card-preview" 
        v-for="card in displayedCards" 
        :key="card.hashid"
        @click="openCard(card)"
      >
        <img :src="card.avatar" :alt="card.name" />
        <div class="card-name">{{ card.name || '未知卡牌' }}</div>
      </div>
    </div>

    <!-- 卡牌详情弹窗 -->
    <CardModal
      :visible="showCardModal"
      :card="selectedCard"
      :showActions="true"
      :cardActions="[
        { label: selectedCard?.onsale ? '下架' : '上架', handler: toggleSale },
        { label: '销毁', handler: destroyCard }
      ]"
      @close="closeCardModal"
    />
  </div>
</template>

<script>
import { request } from "../api/request"
import CardModal from "../components/CardModal.vue"

export default {
  components: { CardModal },
  data() {
    return {
      cards: [],
      displayedCards: [],
      sortKey: "rarity",
      filteronsale: false,
      showCardModal: false,
      selectedCard: null
    }
  },
  async mounted() {
    await this.fetchCards()
  },
  methods: {
    async fetchCards() {
      try {
        const res = await request("/auth/card/query", { method: "POST" })
        if (res.code === 200) {
          this.cards = res.data
          this.applyFilterAndSort()
        } else {
          this.$root.showAlert("danger", "获取收藏卡牌失败");
        }
      } catch (e) {
        console.error(e)
      }
    },
    applyFilterAndSort() {
      let temp = [...this.cards]
      if (this.filteronsale) temp = temp.filter(c => c.onsale)
      if (this.sortKey === "name") temp.sort((a,b) => (a.name || '').localeCompare(b.name || ''))
      else if (this.sortKey === "rarity") {
        const rarityOrder = { "common":1, "rare":2, "epic":3, "legendary":4 }
        temp.sort((a,b) => (rarityOrder[b.rarity]||0) - (rarityOrder[a.rarity]||0))
      }
      this.displayedCards = temp
    },
    openCard(card) {
      this.selectedCard = card
      this.showCardModal = true
    },
    closeCardModal() {
      this.showCardModal = false
      this.selectedCard = null
    },
    async destroyCard(card) {
      if (!confirm("确定要销毁这张卡牌吗？")) return
      try {
        const res = await request("/auth/card/destroy", {
          method: "POST",
          body: JSON.stringify({ hashid: card.hashid })
        })
        if (res.code === 200) {
          this.$root.showAlert("success", "销毁成功");
          await this.fetchCards()
          this.closeCardModal()
        } else alert(res.msg)
      } catch (e) {
        console.error(e)
        this.$root.showAlert("danger", "销毁失败");
      }
    },
    async toggleSale(card) {
  try {
    if (!card.onsale) {
      // 上架前弹出输入框设置价格
      let inputPrice = prompt("请输入上架价格（GoCoin）", "10");
      if (inputPrice === null) return; // 用户取消
      inputPrice = parseInt(inputPrice);
      if (isNaN(inputPrice) || inputPrice <= 0) {
        this.$root.showAlert("warning", "请输入有效数字");
        return;
      }
      // 发送上架请求
      const res = await request("/auth/card/sell", {
        method: "POST",
        body: JSON.stringify({ hashid: card.hashid, cost: inputPrice })
      });
      if (res.code === 200) {
        this.$root.showAlert("success", "上架成功");
        card.onsale = true;        // 更新当前卡牌状态
        card.price = inputPrice;    // 保存上架价格
        this.fetchCards();
      } else {
        this.$root.showAlert("danger", "上架失败");
      }
    } else {
      // 下架逻辑
      const res = await request("/auth/card/cancel", {
        method: "POST",
        body: JSON.stringify({ hashid: card.hashid })
      });
      if (res.code === 200) {
        this.$root.showAlert("success", "下架成功");
        card.onsale = false;
        this.fetchCards();
      } else {
        this.$root.showAlert("danger", "下架失败");
      }
    }
  } catch (e) {
    console.error(e);
    this.$root.showAlert("danger", "操作失败");
  }
}

  }
}
</script>

<style scoped>
.collection-page { padding: 1rem; max-width:900px; margin:0 auto; }
h2 { margin-bottom:1rem; color:#1e3a8a; }
.controls { display:flex; gap:2rem; margin-bottom:1rem; align-items:center; }
.card-grid { display:grid; grid-template-columns:repeat(auto-fill,minmax(150px,1fr)); gap:1rem; }
.card-preview { position: relative;width:150px; height:200px; overflow:hidden; border-radius:8px; cursor:pointer; display:flex; flex-direction:column; justify-content:flex-end; align-items:center; background:#f5f5f5; transition:transform 0.2s; }
.card-preview img { width:100%; height:100%; object-fit:cover; }
.card-preview:hover { transform:scale(1.05); }
.card-name { width:100%; text-align:center; background:rgba(0,0,0,0.4); color:white; font-weight:bold; padding:0.2rem 0; position:absolute; bottom:0; left:0; border-radius:0 0 8px 8px; }
</style>
