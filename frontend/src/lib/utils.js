import { clsx } from "clsx"
import { twMerge } from "tailwind-merge"


export function cn(...inputs) {
  return twMerge(clsx(inputs));
}

// JS 里没有类型系统，所以不需要 ObjectValues<T> 这种定义
// 如果真的需要类似功能，可以写一个工具函数
// 比如获取对象的值数组：
export function objectValues(obj) {
  return Object.keys(obj).map(key => obj[key]);
}


