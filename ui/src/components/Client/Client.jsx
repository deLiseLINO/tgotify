import React from "react";
import styled from "styled-components";
import bin from "../../img/bin.svg";
import pen from "../../img/pen.svg";


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

const Icon = styled.img`
  cursor: pointer;
`;


const Client = ({props}) => {

  return (
    <ClientPerson>
      <ClientName>{props.name}</ClientName>
      <ClientToken>{props.token}</ClientToken>
      <ClientIcons>
        <Icon src={pen} alt="pen" />
        <Icon src={bin} alt="bin" />
      </ClientIcons>
    </ClientPerson>
  );
};

export default Client;
