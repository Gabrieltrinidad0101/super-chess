import React, { useEffect, useState, useContext } from 'react'
import IUser, { IUserState } from '../types/user'
import { Outlet, useNavigate } from 'react-router-dom'
import { customFetch } from '../utils/customFetch'

const userInitialState = {
  name: '',
  password: '',
  _id: ''
}

const AuthContext = React.createContext<IUserState>({
  user: userInitialState,
  setUser: (user: IUser) => user
})

const AuthenticationProvider = (): JSX.Element => {
  const [user, setUser] = useState<IUser>(userInitialState)
  const navigation = useNavigate()

  const verifyAuthentication = async (): Promise<boolean> => {
    try {
      const [error,user] = await customFetch.get<IUser>('/verifyAuthentication')
      if (error) return true
      setUser(user)
      return false
    } catch (error) {
      console.log(error)
    }
    return true
  }

  useEffect(() => {
    verifyAuthentication()
      .then((noHasAccount) => {
        if (noHasAccount) navigation('/register')
      })
      .catch(error => {
        console.log(error)
      })
  }, [])

  const containerSetUser = (user: IUser): void => {
    setUser((prevUser: IUser | undefined) => ({ ...prevUser, ...user }))
  }
  return <AuthContext.Provider value={{ user, setUser: containerSetUser }} >{
    user?._id ? <>Loading</> : <Outlet />
  }</AuthContext.Provider>
}
export const useAuthenticationContext = () => {
  return useContext(AuthContext)
}

export { AuthContext, AuthenticationProvider }
