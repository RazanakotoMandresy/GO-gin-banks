import React, { useState } from "react";
import { FiSettings } from "react-icons/fi";
import "./HomeHeader.css";
import { useAppContext } from "../../App";
import { GetDate } from "../../logics/funLogic/func";

const HomeHeader = () => {
  const { userData } = useAppContext();
  const [actualDate, _] = useState(GetDate());
  return (
    <div className="head">
      <ul>
        <li className="left"> Bonjour {userData.name}</li>
        <li className="right">{actualDate}</li>
        <li className="right">
          <FiSettings /> Paramètres globaux
        </li>
      </ul>
    </div>
  );
};

export default HomeHeader;
