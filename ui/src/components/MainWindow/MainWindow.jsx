import React, { useState, useEffect } from "react";
import axios from "../../utils/axios";
import styled from "styled-components";
import person from "../../img/person.svg";
import LogOut from "../../img/LogOut.svg";
import Client from "../Client/Client";
import bin from "../../img/bin.svg";
import pen from "../../img/pen.svg";

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

const ClientPerson = styled.div`
  width: 572px;
  min-height: 74px;

  gap: 70px;

  justify-content: space-between;

  margin-bottom: 20px;
  border-radius: 10px;
  background: #699bf72d;

  padding-left: 20px;
  padding-right: 20px;

  display: flex;
  flex-direction: row;
  align-items: center;
  box-shadow: 4px 4px 10px 0px #f1f6fc;
  backdrop-filter: blur(6.449999809265137px);
`;

const ClientName = styled.div`
  color: #699bf7;
`;

const ClientToken = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;

  color: #fff;
  border-radius: 6px;

  width: 100%;
  height: 28px;
  background-color: #699bf7;

  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 400px;
`;

const ClientIcons = styled.div`
  display: flex;
  gap: 20px;
`;

const ButtonPlace = styled.div`
  display: flex;
  justify-content: space-between;
  width: 600px;
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

const MainWindow = (props) => {
  const [arrUsers, setArrUsers] = useState([]);


  const openPersonalDate = () => {
    console.log(1);
  }

  const moveLogOut = () => {
    localStorage.removeItem("token");
    console.log('удлили токен, дальше удачи');
    // props.isAuth(false)
  }

  const sendMessage = () => {
    let token = localStorage.getItem("token");

    const data = {
      text: "тестовое сообщение",
    };

    const config = {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    };

    axios
      .post("/message", data, config)
      .then((response) => {
        console.log(response);
        console.log('Сообщение отправлено');
      })
      .catch((err) => {
        console.error(err);
      });
  };

  useEffect(() => {
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
  }, []);

  return (
    <MainBackground>
      <MainSettings />
      <MainSettingsIcons>
        <Icon onClick={openPersonalDate} src={person} alt="person" />
        <Icon onClick={moveLogOut} src={LogOut} alt="LogOut" />
      </MainSettingsIcons>

      <ButtonPlace>
        <ButtonSlyled onClick={sendMessage}>sendMessage</ButtonSlyled>
        <ButtonSlyled>Add Client</ButtonSlyled>
      </ButtonPlace>

      <ClientsPlace>
        <ClientPerson>
          <ClientName>Name</ClientName>
          <ClientToken>*********</ClientToken>
          <ClientIcons>
            <Icon src={pen} alt="pen" />
            <Icon src={bin} alt="bin" />
          </ClientIcons>
        </ClientPerson>

        {arrUsers.map((item) => (
          <Client props={item} />
        ))}

        {/* <Client>
          <ClientName>Name</ClientName>
          <ClientToken>*********</ClientToken>
          <ClientIcons>
            <Icon src={pen} alt="pen" />
            <Icon src={bin} alt="bin" />
          </ClientIcons>
        </Client> */}
      </ClientsPlace>
    </MainBackground>
  );
};

export default MainWindow;
