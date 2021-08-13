import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';

const useStyles = makeStyles((theme) => ({
  formControl: {
    margin: theme.spacing(0),
    minWidth: 265,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
}));

const PaymentForm = ({ paymentValues, changePaymentValue, coo, goo, creditCardType, setCreditCardType }) => {


  const classes = useStyles();

  const handleChange = (event) => {
    const name = event.target.name;
    setCreditCardType({
      ...creditCardType,
      [name]: event.target.value,
    });
  };

  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom>
        Credit Card details
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} md={6}>
          <TextField
            required
            id="cardName"
            label="Name on card"
            fullWidth
            value={paymentValues.cardName}
            onChange={(e) => changePaymentValue('cardName', e.target.value)}
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <FormControl className={classes.formControl}>
            <InputLabel htmlFor="cardType">Credit Card Type</InputLabel>
            <Select
              native
              value={creditCardType['cardType']}
              onChange={handleChange}
              name="cardType"
              inputProps={{
                id: 'cardType',
              }}
            >
              <option aria-label="None" value="" />
              <option value="Visa">Visa</option>
              <option value="Master Card">Master Card</option>
              <option value="American Express">American Express</option>
              <option value="Discover">Discover</option>
            </Select>
          </FormControl>
        </Grid>
        <Grid item xs={12} md={6}>
          <TextField
            required
            id="branch"
            label="Institution/Branch"
            fullWidth
            value={paymentValues.branch}
            onChange={(e) => changePaymentValue('branch', e.target.value)}
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <TextField
            required
            id="cardBranding"
            label="Credit Card Branding"
            fullWidth
            value={paymentValues.cardBranding}
            onChange={(e) => changePaymentValue('cardBranding', e.target.value)}
          />
        </Grid>
        <Grid item xs={12}>
          <FormControlLabel
            control={<Checkbox color="secondary" name="saveCard" value="yes" />}
            label="Remember credit card details for next time"
          />
        </Grid>
      </Grid>
    </React.Fragment>
  );
}
export default PaymentForm;