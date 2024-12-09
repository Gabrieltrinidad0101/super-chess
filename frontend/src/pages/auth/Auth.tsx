import React, { useState } from 'react'
import logo from '../../assets/logo.png'
import './AuthComponent.css'
import { Link, useNavigate } from 'react-router-dom'
import IUser from '../../types/user'
import { useAuthenticationContext } from '../../context/AuthenticationProvider'
import { customFetch } from '../../utils/customFetch'

export default function AuthComponent ({isLogin}: {isLogin: boolean}): JSX.Element {
  const [user, setUser] = useState<IUser>({
    name: '',
    password: '',
    isRegister: isLogin
  })

  const navigation = useNavigate()
  const userState = useAuthenticationContext()

  const clickAuth = async () => {
    const [error,response] = await customFetch.post('/login',user)
    if(error) return
    console.log(response)
  }

  const inputChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    const { name, value } = e.target
    setUser({ ...user, [name]: value })
  }

  return (
    <div className='login-screen'>
      <div className="screen-1">
        <div className="logo">
          <img alt="chess super" src={logo} />
        </div>
        <div className="input-auth mb-3">
          <label htmlFor="email">Name</label>
          <div className="sec-2">
            <input type="text" name="name" onChange={inputChange} placeholder="Username" />
          </div>
        </div>
        <div className="input-auth">
          <label htmlFor="password">Password</label>
          <div className="sec-2">
            <input type="password" onChange={inputChange} name="password" placeholder="············" />
            <i className="show-hide"></i>
          </div>
        </div>
        <button id="auth-button" className="login" onClick={clickAuth}>{!isLogin ? 'Login' : 'Register'}</button>
        <Link to={!isLogin ? '/login' : '/register'}>
          {isLogin ? 'Login' : 'Register'}
        </Link>
      </div>
    </div>
  )
}