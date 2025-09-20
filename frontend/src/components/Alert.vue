<template>
  <div class="wrapper">
    <div
      v-for="(alert, index) in alerts"
      :key="index"
      :class="['alert', `alert--${alert.type}`]"
    >
      <svg class="alert__icon" width="64" height="64" viewBox="0 0 64 64" fill="none">
        <circle class="alert__icon-path" cx="32" cy="32" r="30.5" stroke="currentColor" />
        <path
          v-if="alert.type === 'success'"
          class="alert__icon-path alert__icon-path--type"
          d="M15 33.5L25 43.5L48.5 20"
          stroke="currentColor"
        />
        <path
          v-else-if="alert.type === 'danger'"
          class="alert__icon-path alert__icon-path--type"
          d="M20 44L44 20 M44 44L20 20"
          stroke="currentColor"
        />
        <path
          v-else-if="alert.type === 'warning'"
          class="alert__icon-path alert__icon-path--type"
          d="M32 19V38 M32 41L32 45"
          stroke="currentColor"
        />
      </svg>
      <div class="alert__message">{{ alert.message }}</div>
      <div class="alert__close" @click="removeAlert(index)">&times;</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Alert",
  data() {
    return {
      alerts: [],
    };
  },
  methods: {
    createAlert(type = "default", message = "") {
      const msg =
        message ||
        (type === "success"
          ? "This is a success notification"
          : type === "danger"
          ? "This is a danger notification"
          : type === "warning"
          ? "This is a warning notification"
          : "This is a default notification");
      this.alerts.push({ type, message: msg });

      // 自动关闭
      setTimeout(() => {
        this.alerts.shift();
      }, 3000);
    },
    removeAlert(index) {
      this.alerts.splice(index, 1);
    },
  },
};
</script>

<style scoped lang="scss">
$blue: #09c8ff;
$red: #ff4555;
$green: #39d7b3;
$yellow: #ffba4b;
$text: #677b7a;

.alert {
  background: white;
  width: 100%;
  display: flex;
  padding: 1rem;
  align-items: center;
  border-radius: 0.8rem;
  border: 2px solid;
  margin-bottom: 1rem;
  position: relative;
  box-shadow: 0px 5px 12px rgba(black, 0.04), 0px 12px 25px rgba(black, 0.07);
  animation: transitionIn 200ms ease forwards;
  transition: all 200ms ease;

  &__message {
    color: $text;
    flex: 1;
  }

  &__close {
    width: 1.3rem;
    height: 1.3rem;
    line-height: 1.3rem;
    color: $text;
    text-align: center;
    position: absolute;
    top: 0.5rem;
    right: 0.5rem;
    cursor: pointer;
    border-radius: 0.4rem;
    &:hover {
      background: rgba($text, 0.2);
    }
  }

  &__icon {
    margin-right: 1rem;
    width: 36px;
    height: 36px;
    animation: fade 500ms linear forwards;
  }

  &__icon-path {
    stroke-dasharray: 1000;
    stroke-dashoffset: 1000;
    stroke-width: 4px;
    animation: dash 1s ease forwards;
    &--type {
      animation-delay: 300ms;
      animation-duration: 0.5s;
    }
  }

  &--success {
    color: $green;
    border-color: $green;
  }
  &--warning {
    color: $yellow;
    border-color: $yellow;
  }
  &--danger {
    color: $red;
    border-color: $red;
  }
}

@keyframes dash {
  from {
    stroke-dashoffset: 1000;
  }
  to {
    stroke-dashoffset: 0;
  }
}

@keyframes fade {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.wrapper {
  position: fixed;
  top: 1rem;             /* 离顶部一点点 */
  left: 50%;
  transform: translateX(-50%); /* 只水平居中 */
  max-width: 500px;
  width: 80%; // 适配小屏幕
  z-index: 9999;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

</style>
