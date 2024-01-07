import React from 'react';
import styled from "styled-components";
import MainWindow from '../MainWindow/MainWindow';


const MainBackground = styled.div`
    width: 1200px;
    height: 100vh;
    margin: 0 auto;

    display: flex;
    align-items: center;
    justify-content: center;
`;


const MainPage = () => {
    return (
        <MainBackground>
            <MainWindow />
        </MainBackground>
    );
};

export default MainPage;