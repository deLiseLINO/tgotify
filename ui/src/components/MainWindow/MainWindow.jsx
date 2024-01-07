import React, { useState, useEffect } from "react";
import axios from "../../utils/axios";
import styled from "styled-components";
import person from "../../img/person.svg";
import LogOut from "../../img/LogOut.svg";
import Client from "../Client/Client";

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
`;

const ButtonSlyled = styled.button`
    cursor: pointer;

    border-radius: 6px;
    background: #699BF7;
    padding: 12px 28px;
    border: 0;

    color: #fff;
    display: flex;
    justify-content: center;
    align-items: center;


    &:hover{
        background: #8ab1fa;
    }
`;





const MainWindow = () => {
  const [arrUsers, setArrUsers] = useState([]);

  const sendMessage = () => {
    let token = localStorage.getItem("token");

    axios
      .get("/message", {
        headers: {
          Authorization: `Bearer ${token}`,
          message: "тестовое сообщение",
          //   pass: user.password,
        },
      })
      .then((response) => {
        setArrUsers(response.data);
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
        <Icon src={person} alt="person" />
        <Icon src={LogOut} alt="LogOut" />
      </MainSettingsIcons>

      <ButtonPlace>
        <ButtonSlyled onClick={sendMessage}>sendMessage</ButtonSlyled>
        <ButtonSlyled>Add Client</ButtonSlyled>
      </ButtonPlace>

      <ClientsPlace>
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
