
export default {
  d() {
    if (process.env.NODE_ENV === 'development') {
      console.log(...arguments)
    }
  },
  i() {
    console.info(...arguments)
  }
}
