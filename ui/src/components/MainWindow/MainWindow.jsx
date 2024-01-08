import React, { useState, useEffect } from "react";
import axios from "../../utils/axios";
import styled from "styled-components";
import person from "../../img/person.svg";
import LogOut from "../../img/LogOut.svg";
import Client from "../Client/Client";
import ClientSample from "../ClientSample/ClientSample";

const MainBackground = styled.div`
  position: relative;

  background-color: #f1f6fc;
  width: 1000px;
  height: 600px;

  display: flex;
  flex-direction: column;

  align-items: center;
  justify-content: center;
`;

const MainSettings = styled.div`
  position: absolute;
  right: 0;
  top: 0;
  width: 156px;
  height: 70px;

  border-radius: 0px 0px 0px 20px;
  background: #fff;
`;

const MainSettingsIcons = styled.div`
  position: absolute;
  right: 26px;
  top: 10px;

  display: flex;
  justify-content: center;
  gap: 30px;
`;

const Icon = styled.img`
  cursor: pointer;

  &:hover{
    opacity: 0.8;
  }
`;

const ClientsPlace = styled.div`
  display: flex;

  align-items: center;
  justify-content: start;

  flex-direction: column;

  padding-top: 20px;
  padding-bottom: 20px;

  width: 680px;
  height: 300px;

  overflow: auto;
  overflow-x: hidden;

  background-color: #ffffff44;

  &::-webkit-scrollbar {
    width: 12px; /* ширина scrollbar */
  }
  &::-webkit-scrollbar-track {
    background: #ffa6000; /* цвет дорожки */
  }
  &::-webkit-scrollbar-thumb {
    background-color: #ffffff; /* цвет плашки */
    border-radius: 20px; /* закругления плашки */
    border: 3px solid #ffa6000; /* padding вокруг плашки */
  }
`;

const ButtonPlace = styled.div`
  display: flex;
  justify-content: space-between;
  width: 600px;

  height: 46px;

  gap: 40px;
`;

const ButtonSlyled = styled.button`
  cursor: pointer;

  border-radius: 6px;
  background: #699bf7;
  padding: 12px 28px;
  border: 0;

  color: #fff;
  display: flex;
  justify-content: center;
  align-items: center;

  &:hover {
    background: #8ab1fa;
  }
`;

const Input = styled.input`
  display: flex;
  flex-direction: row;

  width: 100%;
  padding-left: 16px;

  border-radius: 8px;
  border: 3px solid #699bf7;

  font-size: 16px;
`



//////


const MainWindow = (props) => {
  const [arrUsers, setArrUsers] = useState([]);
  const [inputValue, setInputValue] = useState('');
  const [itemUpdate, setItemUpdate] = useState(false);
  const [tooglePersonDateValue, setTooglePersonDateValue] = useState(false);


  const handleInputChange = (event) => {
    setInputValue(event.target.value);
  };


  const peopleSetting = () => {
    console.log(tooglePersonDateValue);
  }

  const tooglePersonDate = () => {
    setTooglePersonDateValue(!tooglePersonDateValue)
  }


  const updatePage = (value) =>{
    setItemUpdate(value)
  }


  const moveLogOut = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("name");
    console.log('удлили токен, дальше удачи');
    // props.isAuth(false)
  }

  const sendMessage = () => {
    let token = localStorage.getItem("token");

    const data = {
      text: inputValue,
    };

    setInputValue('')

    const config = {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    };

    axios
      .post("/message", data, config)
      .then((response) => {
        console.log(response.data.status);
      })
      .catch((err) => {
        console.error(err);
      });
  };

  useEffect(() => {
    console.log('обновление данных');

    let token = localStorage.getItem("token");

    axios
      .get("/client", {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((response) => {
        setArrUsers(response.data);
      })
      .catch((err) => {
        console.error(err);
      });
  }, [itemUpdate]);

  return (
    <MainBackground>
      <MainSettings />
      <MainSettingsIcons>
        <Icon onClick={peopleSetting} src={person} alt="person" />
        <Icon onClick={moveLogOut} src={LogOut} alt="LogOut" />
      </MainSettingsIcons>

      <ButtonPlace>
        <ButtonSlyled onClick={sendMessage}>sendMessage</ButtonSlyled>
        <Input type="text" value={inputValue} onChange={handleInputChange} placeholder="message..."/>
        <ButtonSlyled onClick={tooglePersonDate}>Add Client</ButtonSlyled>
      </ButtonPlace>

      <ClientsPlace>
        {tooglePersonDateValue ? <ClientSample update={updatePage}/> : null}
        {arrUsers.map((item) => (
          <Client props={item} />
        ))}
      </ClientsPlace>

    </MainBackground>
  );
};

export default MainWindow;
