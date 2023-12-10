// 封装本地token的读取和写入方法
// export是分别暴露，需要一个个引入
export const SET_TOKEN = (token) =>{
  localStorage.setItem("TOKEN",token)
}

// 我是觉得，存在cookies中，它就会自己附在请求头上，然后用每次请求的请求头去检查比较合理
export const GET_TOKEN = ()=>{
  return localStorage.getItem("TOKEN")
}
