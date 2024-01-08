import React, { useState } from "react";
import styled from "styled-components";
import axios from "../../utils/axios";


const AuthContainer = styled.div`
  display: flex;

  justify-content: center;
  align-items: center;

  height: 600px;
  gap: 0;


  @media (max-width: 768px) {
    grid-template-columns: 1fr;
  }
`;

const AuthContent = styled.div`
  width: 400px;
  height: 460px;
  border-radius: 10px;
  box-shadow: 0px 4px 20px 0px rgba(0, 0, 0, 0.10);

  display: flex;
  flex-direction: column;

  align-items: center;
  justify-content: center;
`;

const AuthForm = styled.form`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  gap: 20px;
`;

const ContainerInput = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 4px;
`;

const AuthImport = styled.input`
  height: 36px;
  padding-left: 10px;
  width: 200px;
  border-radius: 10px;

  border: 1px solid #2600ff;
  background-color: #fff;
`;

const AuthButton = styled.input`
  cursor: pointer;

  margin: 20px;
  height: 36px;
  padding-left: 10px;
  width: 200px;
  border-radius: 10px;

  color: #fff;

  border: 0;
  background-color: #699bf7;

  &:hover {
    background-color: #85affe;
  }
`;

const RadioContainer = styled.div`
  margin-top: 30px;
  margin-bottom: 10px;
  display: flex;
  gap: 20px;
`;

const RadioInput = styled.input`
  position: fixed;
  opacity: 0;
  pointer-events: none;
`;

const RadioLable = styled.label`
  cursor: pointer;

  height: 42px;
  width: 100%;
  padding: 0 20px;
  display: flex;
  justify-content: center;
  align-items: center;

  &:hover {
    background-color: #00c3ff55;
  }

  ${RadioInput}:checked {
    background-color: #ff2a00;
  }
`;

const AuthorizationPage = ({isAuth}) => {

  const [user, setUser] = useState({name:'', password: ''});
  const [isRegistration, setIsRegistration] = useState(false);

  const handleRadioChange = () => {
    setIsRegistration(!isRegistration);
  };

  const handleChange = e => {
    const { name, value } = e.target;
    setUser(prevUser => ({
        ...prevUser,
        [name]: value
    }));
};


  const btnRegistration = (event) => {
    event.preventDefault()

    axios
    .post("/user", {
      "name": user.name,
      "pass": user.password,
    })
    .then((response) => {
      localStorage.setItem('token', response.data.token.toString())
      localStorage.setItem('name', user.name)
      console.log('Registration');
      isAuth(true)
    })
    .catch((err) => {console.error(err)})
  };


  const btnAuthorization = (event) => {
    event.preventDefault()

    axios
    .post("/user/signin", {
      "name": user.name,
      "pass": user.password,
    })
    .then((response) => {
      localStorage.setItem('token', response.data.token.toString())
      console.log('Authorization');
      isAuth(true)
    })
    .catch((err) => {console.error(err)})
  };

  return (
    <AuthContainer>
      <AuthContent>
        <h2>{isRegistration ? "Registration" : "Login"}</h2>

        <AuthForm>
          <RadioContainer>
            <div>
              <RadioLable>
                <RadioInput
                  type="radio"
                  value="registration"
                  checked={isRegistration}
                  onChange={handleRadioChange}
                />
                Registration
              </RadioLable>
            </div>
            <div>
              <RadioLable>
                <RadioInput
                  type="radio"
                  value="authentication"
                  checked={!isRegistration}
                  onChange={handleRadioChange}
                />
                Authentication
              </RadioLable>
            </div>
          </RadioContainer>

          <ContainerInput>
            <label for="name">Login</label>
            <AuthImport 
              onChange={handleChange}
              type="name" 
              id="name" 
              name="name" 
              required />
          </ContainerInput>

          <ContainerInput>
            <label for="password">Password</label>
            <AuthImport
              onChange={handleChange}
              type="password"
              id="password"
              name="password"
              required
            />
          </ContainerInput>

          <AuthButton
            type="submit"
            value={isRegistration ? "Submit" : "Submit"}
            onClick={isRegistration ? btnRegistration : btnAuthorization}
          />
        </AuthForm>
      </AuthContent>
    </AuthContainer>
  );
};

export default AuthorizationPage;
