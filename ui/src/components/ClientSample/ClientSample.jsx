import React, { useState } from "react";
import styled from "styled-components";
import axios from "../../utils/axios";
import btnAdd from "../../img/btnAdd.svg";

const Icon = styled.img`
  cursor: pointer;

  &:hover {
    opacity: 0.8;
  }
`;

const ClientPerson = styled.div`
  width: 572px;
  min-height: 74px;

  gap: 20px;

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

const InputClientName = styled.input`
  color: #699bf7;

  border-radius: 10px;

  height: 28px;
  font-size: 12px;
  border: 0px;
  padding-left: 12px;
  background-color: #ffffff;

  width: 140px;
`;

const InputClientToken = styled.input`
  display: flex;
  align-items: center;
  justify-content: center;

  border-radius: 10px;
  border: 0px;
  padding-left: 12px;

  color: #699bf7;
  border-radius: 6px;

  font-size: 12px;

  width: 100%;
  height: 28px;
  background-color: #ffffff;

  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 400px;
`;

const ClientIcons = styled.div`
  margin-left: 100px;

  display: flex;
  gap: 20px;
`;

const ClientSample = ({update}) => {
  const [user, setUser] = useState({ name: "", token: "" });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setUser((prevUser) => ({
      ...prevUser,
      [name]: value,
    }));
  };

  const addNewUser = () => {

    let token = localStorage.getItem("token");

    const data = {
      name: user.name,
      token: user.token,
    };

    setUser('')

    const config = {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    };

    axios
      .post("/client", data, config)
      .then((response) => {
        console.log(response);
        update(true)
      })
      .catch((err) => {
        console.error(err);
      });

  }



  return (
    <ClientPerson>
      <InputClientName
        onChange={handleChange}
        type="name"
        id="name"
        name="name"
        required
        placeholder="Name"
      />
      <InputClientToken
        onChange={handleChange}
        type="token"
        id="token"
        name="token"
        required
        placeholder="Token"
      />
      <ClientIcons>
        <Icon onClick={addNewUser} src={btnAdd} alt="add" />
      </ClientIcons>
    </ClientPerson>
  );
};

export default ClientSample;
