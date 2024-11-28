export default function Auth({isLogin}: {isLogin: boolean}) {
  return (
    <h1>{isLogin ? "login" : "register"}</h1>
  )
}
