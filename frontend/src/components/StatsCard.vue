<template>
  <div class="stats-card" :class="{ 'stats-card--highlight': highlight }">
    <div class="stats-label">{{ label }}</div>
    <div class="stats-value">{{ value }}</div>
    <div class="stats-trend" v-if="trend">
      <span :class="trend > 0 ? 'up' : 'down'">{{ trend > 0 ? '↑' : '↓' }} {{ Math.abs(trend) }}%</span>
      <span class="trend-label">{{ trendLabel }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  label: string
  value: string | number
  highlight?: boolean
  trend?: number
  trendLabel?: string
}>()
</script>

<style scoped>
.stats-card {
  background: white;
  border: 1px solid #e8e8e8;
  border-radius: 12px;
  padding: 20px 24px;
  transition: transform 0.2s, box-shadow 0.2s;
}

.stats-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(102, 126, 234, 0.1);
}

.stats-card--highlight {
  background: linear-gradient(135deg, #667eea, #764ba2);
  border: none;
  color: white;
}

.stats-card--highlight .stats-label {
  color: rgba(255, 255, 255, 0.8);
}

.stats-card--highlight .stats-value {
  color: white;
}

.stats-label {
  font-size: 13px;
  color: #999;
  margin-bottom: 6px;
}

.stats-value {
  font-size: 28px;
  font-weight: 700;
  color: #333;
}

.stats-trend {
  margin-top: 8px;
  font-size: 12px;
}

.up {
  color: #52c41a;
}

.down {
  color: #ff4d4f;
}

.trend-label {
  color: #999;
  margin-left: 4px;
}

.stats-card--highlight .trend-label {
  color: rgba(255, 255, 255, 0.7);
}
</style>
