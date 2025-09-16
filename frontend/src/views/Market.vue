<template>
  <div class="market-page">
    <h2>市场</h2>

    <div class="balance">当前余额: {{ balance }} GoCoin</div>

    <div class="controls">
      <label>
        搜索：
        <input v-model="searchKey" @input="filterCards" placeholder="输入名称、HashID、稀有度" />
      </label>
      <label>
        排序：
        <select v-model="sortKey" @change="sortCards">
          <option value="rarity">按稀有度</option>
          <option value="name">按名称</option>
        </select>
      </label>
    </div>

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

   <CardModal
  :visible="showCardModal"
  :card="selectedCard"
  :showActions="true"
  :cardActions="[
    { label: '购买 ' + (selectedCard?.price || 0) + ' GoCoin', handler: buyCard }
  ]"
  @close="closeCardModal"
/>
  </div>
</template>

<script>
import { request } from "../api/request"
import CardModal from "../components/CardModal.vue"

export default {
  name: "Market",
  components: { CardModal },
  data() {
    return {
      balance: 0,
      cards: [],
      displayedCards: [],
      sortKey: "rarity",
      searchKey: "",
      showCardModal: false,
      selectedCard: null
    }
  },
  async mounted() {
    await this.fetchBalance()
    await this.fetchMarketCards()
  },
  methods: {
    async fetchBalance() {
      try {
        const res = await request("/auth/user/profile", { method:"POST", body: "{}" })
        if(res.code===200) this.balance = res.data.gocoin
      } catch(e){console.error(e)}
    },
    async fetchMarketCards() {
      try {
        const res = await request("/auth/market/query",{ method:"POST" })
        if(res.code===200){
          const marketCards = await Promise.all(res.data.map(async t=>{
            const cardRes = await request("/auth/card/query",{ method:"POST", body: JSON.stringify({ hashid: t.CardID }) })
            const cardInfo = cardRes.data[0] || {}
            return { ...cardInfo, price:t.Price, seller:t.Seller, transID:t.TransID || t.transID }
          }))
          this.cards = marketCards
          this.displayedCards = [...this.cards]
          this.sortCards()
        }
      } catch(e){console.error(e)}
    },
    sortCards(){
      if(this.sortKey==="name") this.displayedCards.sort((a,b)=>(a.name||'').localeCompare(b.name||''))
      else if(this.sortKey==="rarity"){
        const rarityOrder={ "普通":1, "稀有":2, "史诗":3, "传说":4 }
        this.displayedCards.sort((a,b)=>(rarityOrder[b.rarity]||0)-(rarityOrder[a.rarity]||0))
      }
    },
    filterCards(){
      const key=this.searchKey.toLowerCase()
      this.displayedCards=this.cards.filter(c=>
        (c.name && c.name.toLowerCase().includes(key)) ||
        (c.hashid && c.hashid.toLowerCase().includes(key)) ||
        (c.rarity && c.rarity.toLowerCase().includes(key))
      )
      this.sortCards()
    },
    openCard(card){
      this.selectedCard = card
      this.showCardModal = true
    },
    closeCardModal(){
      this.showCardModal = false
      this.selectedCard = null
    },
    async buyCard(card){
      if(this.balance<card.price){ this.$root.showAlert("warning", "资金不足"); return }
      if(!confirm(`确认购买 ${card.name} ？`)) return
      try{
        const res=await request("/auth/market/buy",{
          method:"POST",
          body: JSON.stringify({ orderid: card.transID })
        })
        if(res.code===200){
          this.$root.showAlert("success", "购买成功");
          window.location.reload();
          this.balance -= card.price
          await this.fetchMarketCards()
          this.closeCardModal()
        } else this.$root.showAlert("danger", "购买失败");
      } catch(e){ console.error(e); this.$root.showAlert("danger", "购买失败"); }
    }
  }
}
</script>

<style scoped>
.market-page { max-width:900px; margin:0 auto; padding:1rem; }
h2 { color:#1e3a8a; margin-bottom:0.5rem; }
.balance { font-weight:bold; margin-bottom:1rem; }
.controls { display:flex; gap:1rem; margin-bottom:1rem; align-items:center; }
.controls input, .controls select { padding:0.3rem; }
.card-grid { display:grid; grid-template-columns:repeat(auto-fill,minmax(150px,1fr)); gap:1rem; }
.card-preview { position: relative;width:150px; height:200px; overflow:hidden; border-radius:8px; cursor:pointer; display:flex; flex-direction:column; justify-content:flex-end; align-items:center; background:#f5f5f5; transition:transform 0.2s; }
.card-preview img { width:100%; height:100%; object-fit:cover; }
.card-preview:hover { transform:scale(1.05); }
.card-name { width:100%; text-align:center; background:rgba(0,0,0,0.4); color:white; font-weight:bold; padding:0.2rem 0; position:absolute; bottom:0; left:0; border-radius:0 0 8px 8px; }
</style>
