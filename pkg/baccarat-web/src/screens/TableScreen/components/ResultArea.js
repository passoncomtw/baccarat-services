import Grid from "@mui/material/Grid";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import Card from "screens/PokerScreen/components/cards/Card";
import HiddenCard from "screens/PokerScreen/components/cards/HiddenCard";

const styles = {
  containerStyle: { height: 450 },
  emptyCardArea: { height: 200, width: 160, backgroundColor: "#CCC" },
  cardStyle: { height: 200, width: 160 },
  stackStyle: { margin: 10 },
  middleStyle: {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  },
  pokerStackStyle: { display: "flex", flex: 1, alignItems: "center" },
  rightAreaStyle: {
    display: "flex",
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
};

const TITLES = {
  Ready: "準備投注",
  Betting: "開始投注",
  StopBetting: "停止下注",
};

const EmptyResult = (props) => {
  const { timeLeft, step } = props;
  return (
    <Grid container>
      <Grid item xs={4} style={styles.containerStyle}>
        <Stack direction="row" spacing={2} style={styles.stackStyle}>
          <div style={styles.emptyCardArea}></div>
          <div style={styles.middleStyle}>
            <Typography variant="h2" component="h2">
              莊
            </Typography>
          </div>
          <div style={styles.emptyCardArea}></div>
        </Stack>
        <Stack direction="row" spacing={2} style={styles.stackStyle}>
          <div style={styles.emptyCardArea}></div>
          <div style={styles.middleStyle}>
            <Typography variant="h2" component="h2">
              閒
            </Typography>
          </div>
          <div style={styles.emptyCardArea}></div>
        </Stack>
      </Grid>
      <div style={styles.pokerStackStyle}>
        <HiddenCard containerStyle={{ width: 100, height: 160 }} />
      </div>
      <Grid item xs={3} flex={2}>
        <div style={styles.rightAreaStyle}>
          <div>
            <Typography variant="h1" component="h1">
              {TITLES[step]}
            </Typography>
            <Typography variant="h2" component="h2" textAlign="center">
              {timeLeft / 1000}
            </Typography>
          </div>
        </div>
      </Grid>
    </Grid>
  );
};

const ResultContent = (props) => {
  const { timeLeft, step } = props;
  return (
    <Grid container>
      <Grid item xs={6} style={styles.containerStyle}>
        <Stack direction="row" spacing={2} style={styles.stackStyle}>
          <Card
            containerStyle={styles.cardStyle}
            cardData={{
              animationDelay: 0,
              cardFace: "1",
              suit: "Spade",
              value: 1,
            }}
          />
          <div style={styles.middleStyle}>
            <Typography variant="h2" component="h2">
              莊
            </Typography>
          </div>
          <Card
            containerStyle={styles.cardStyle}
            cardData={{
              animationDelay: 0,
              cardFace: "2",
              suit: "Spade",
              value: 2,
            }}
          />
          <Card
            containerStyle={styles.cardStyle}
            cardData={{
              animationDelay: 0,
              cardFace: "3",
              suit: "Spade",
              value: 3,
            }}
          />
        </Stack>
        <Stack direction="row" spacing={2} style={styles.stackStyle}>
        <Card
            containerStyle={styles.cardStyle}
            cardData={{
              animationDelay: 0,
              cardFace: "4",
              suit: "Spade",
              value: 4,
            }}
          />
          <div style={styles.middleStyle}>
            <Typography variant="h2" component="h2">
              閒
            </Typography>
          </div>
          <Card
            containerStyle={styles.cardStyle}
            cardData={{
              animationDelay: 0,
              cardFace: "5",
              suit: "Spade",
              value: 5,
            }}
          />
          <Card
            containerStyle={styles.cardStyle}
            cardData={{
              animationDelay: 0,
              cardFace: "6",
              suit: "Spade",
              value: 6,
            }}
          />
        </Stack>
      </Grid>
      <div style={styles.pokerStackStyle}>
        <HiddenCard containerStyle={{ width: 100, height: 160 }} />
      </div>
      <Grid item xs={3} flex={2}>
        <div style={styles.rightAreaStyle}>
          <div>
            <Typography variant="h1" component="h1">
              {TITLES[step]}
            </Typography>
            <Typography variant="h2" component="h2" textAlign="center">
              {timeLeft / 1000}
            </Typography>
          </div>
        </div>
      </Grid>
    </Grid>
  );
};

const ResultArea = (props) => {
  const { timeLeft, step } = props;
  if (step === "StopBetting") {
    return <ResultContent timeLeft={timeLeft} step={step} />;
  }

  return (<EmptyResult timeLeft={timeLeft} step={step} />);
};

export default ResultArea;
