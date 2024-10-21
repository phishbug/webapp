import type { FunctionComponent } from "react";
import ListPost from '../components/Posts/ListPost'
import { React, useEffect, useState } from 'react';
import axios from 'axios';

export default function Dashboard() {
  return (
    <div id="dashboard">
      <ListPost/>
     
    </div>
  );
}
