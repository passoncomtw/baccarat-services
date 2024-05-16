import Grid from "@mui/material/Grid";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import HiddenCard from "screens/PokerScreen/components/cards/HiddenCard";

const styles = {
  containerStyle: { height: 450 },
  emptyCardArea: { height: 200, width: 160, backgroundColor: "#CCC" },
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

const ResultArea = (props) => {
  const { timeLeft } = props;
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
              準備投注
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

export default ResultArea;
