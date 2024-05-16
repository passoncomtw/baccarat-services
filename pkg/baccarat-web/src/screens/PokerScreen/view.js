import "raf/polyfill";

import React, { Component } from "react";
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import Grid from "@mui/material/Grid";
import Spinner from "./Spinner";
import WinScreen from "./WinScreen";

import Card from "./components/cards/Card";
import HiddenCard from "./components/cards/HiddenCard";

class PokerScreen extends Component {
  render() {
    return (
      <div>
        <Grid container>
          <Grid item xs={1}>
            <HiddenCard />
          </Grid>
        </Grid>
        <Grid container>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "1",
                suit: "Spade",
                value: 1,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "2",
                suit: "Spade",
                value: 2,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "3",
                suit: "Spade",
                value: 3,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "4",
                suit: "Spade",
                value: 4,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "5",
                suit: "Spade",
                value: 5,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "6",
                suit: "Spade",
                value: 6,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "7",
                suit: "Spade",
                value: 7,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "8",
                suit: "Spade",
                value: 8,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "9",
                suit: "Spade",
                value: 9,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "9",
                suit: "Spade",
                value: 9,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "10",
                suit: "Spade",
                value: 10,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "11",
                suit: "Spade",
                value: 11,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "12",
                suit: "Spade",
                value: 12,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "13",
                suit: "Spade",
                value: 13,
              }}
            />
          </Grid>
        </Grid>

        <Grid container>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "1",
                suit: "Club",
                value: 1,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "2",
                suit: "Club",
                value: 2,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "3",
                suit: "Club",
                value: 3,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "4",
                suit: "Club",
                value: 4,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "5",
                suit: "Club",
                value: 5,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "6",
                suit: "Club",
                value: 6,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "7",
                suit: "Club",
                value: 7,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "8",
                suit: "Club",
                value: 8,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "9",
                suit: "Club",
                value: 9,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "9",
                suit: "Club",
                value: 9,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "10",
                suit: "Club",
                value: 10,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "11",
                suit: "Club",
                value: 11,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "12",
                suit: "Club",
                value: 12,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "13",
                suit: "Club",
                value: 13,
              }}
            />
          </Grid>
        </Grid>
        <Grid container>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "1",
                suit: "Heart",
                value: 1,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "2",
                suit: "Heart",
                value: 2,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "3",
                suit: "Heart",
                value: 3,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "4",
                suit: "Heart",
                value: 4,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "5",
                suit: "Heart",
                value: 5,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "6",
                suit: "Heart",
                value: 6,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "7",
                suit: "Heart",
                value: 7,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "8",
                suit: "Heart",
                value: 8,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "9",
                suit: "Heart",
                value: 9,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "9",
                suit: "Heart",
                value: 9,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "10",
                suit: "Heart",
                value: 10,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "11",
                suit: "Heart",
                value: 11,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "12",
                suit: "Heart",
                value: 12,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "13",
                suit: "Heart",
                value: 13,
              }}
            />
          </Grid>
        </Grid>
        <Grid container>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "1",
                suit: "Diamond",
                value: 1,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "2",
                suit: "Diamond",
                value: 2,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "3",
                suit: "Diamond",
                value: 3,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "4",
                suit: "Diamond",
                value: 4,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "5",
                suit: "Diamond",
                value: 5,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "6",
                suit: "Diamond",
                value: 6,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "7",
                suit: "Diamond",
                value: 7,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "8",
                suit: "Diamond",
                value: 8,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "9",
                suit: "Diamond",
                value: 9,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "9",
                suit: "Diamond",
                value: 9,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "10",
                suit: "Diamond",
                value: 10,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "11",
                suit: "Diamond",
                value: 11,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "12",
                suit: "Diamond",
                value: 12,
              }}
            />
          </Grid>
          <Grid item xs={1}>
            <Card
              cardData={{
                animationDelay: 0,
                cardFace: "13",
                suit: "Diamond",
                value: 13,
              }}
            />
          </Grid>
        </Grid>

        <img alt="test1" style={{width: 50, height: 50}} src={require("./assets/dollar.png")} />
        <img alt="test2" style={{width: 50, height: 50}} src={'/assets/chip.svg'} />
        <Stack direction="row" spacing={2}>
          <Button variant="contained" color="success" style={{color: "#FFFFFF", width: 60, height: 60, borderRadius: 60}}>
            25
          </Button>
          <Button variant="contained" color="warning" style={{color: "#FFFFFF", width: 60, height: 60, borderRadius: 60}}>
            50
          </Button>
          <Button variant="contained" color="error" style={{color: "#FFFFFF", width: 60, height: 60, borderRadius: 60}}>
            100
          </Button>
          <Button variant="contained" color="primary" style={{color: "#FFFFFF", width: 60, height: 60, borderRadius: 60}}>
            500
          </Button>
          <Button variant="contained" color="secondary" style={{color: "#FFFFFF", width: 60, height: 60, borderRadius: 60}}>
            1000
          </Button>
          <Button variant="contained" color="info" style={{color: "#FFFFFF", width: 60, height: 60, borderRadius: 60}}>
            5000
          </Button>
        </Stack>
        <Spinner />
        <WinScreen />
      </div>
    );
  }
}

export default PokerScreen;
