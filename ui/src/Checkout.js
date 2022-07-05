import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Paper from '@material-ui/core/Paper';
import Stepper from '@material-ui/core/Stepper';
import Step from '@material-ui/core/Step';
import StepLabel from '@material-ui/core/StepLabel';
import Button from '@material-ui/core/Button';
import Link from '@material-ui/core/Link';
import Typography from '@material-ui/core/Typography';
import ApplicantForm from './ApplicantForm';
import CreditCardForm from './CreditCardForm';
import Review from './Review';
import ResidencyForm from './ResidencyForm';
import axios from "axios";

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://github.com/flakrimjusufi">
        Flakrim
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles((theme) => ({
  appBar: {
    position: 'relative',
  },
  layout: {
    width: 'auto',
    marginLeft: theme.spacing(2),
    marginRight: theme.spacing(2),
    [theme.breakpoints.up(600 + theme.spacing(2) * 2)]: {
      width: 600,
      marginLeft: 'auto',
      marginRight: 'auto',
    },
  },
  paper: {
    marginTop: theme.spacing(3),
    marginBottom: theme.spacing(3),
    padding: theme.spacing(2),
    [theme.breakpoints.up(600 + theme.spacing(3) * 2)]: {
      marginTop: theme.spacing(6),
      marginBottom: theme.spacing(6),
      padding: theme.spacing(3),
    },
  },
  stepper: {
    padding: theme.spacing(3, 0, 5),
  },
  buttons: {
    display: 'flex',
    justifyContent: 'flex-end',
  },
  button: {
    marginTop: theme.spacing(3),
    marginLeft: theme.spacing(1),
  },
}));

const steps = ['Personal', 'Residency', 'Credit Card', 'Review'];

function getStepContent(step, formValues = null, changeFormValue = null, date = null, setSelectedDate = null, type, setType) {
  switch (step) {
    case 0:
      return <ApplicantForm applicantFormValues={formValues} changeApplicantFormValues={changeFormValue} dateOfBirth={date} setDateOfBirth={setSelectedDate} employment={type} setEmploymentType={setType} />;
    case 1:
      return <ResidencyForm ResidencyFormValues={formValues} changeResidencyFormValues={changeFormValue} />;
    case 2:
      return <CreditCardForm paymentValues={formValues} changePaymentValue={changeFormValue} date setSelectedDate creditCardType={type} setCreditCardType={setType} />;
    case 3:
      return <Review reviewValues={formValues} />;
    default:
      throw new Error('Unknown step');
  }
}

export default function Checkout() {
  const classes = useStyles();
  const [activeStep, setActiveStep] = React.useState(0);
  let formValues = {
    firstName: '',
    lastName: '',
    dateOfBirth: '',
    phoneNumber: '',
    socialSecurityNumber: '',
    employmentType: '',
    occupation: '',
    monthlyIncome: '',
    yearsEmployed: '',
    streetAddress: '',
    yearsAtAddress: '',
    city: '',
    state: '',
    zip: '',
    country: '',
    ownership: '',
    monthlyPayment: '',
    cardName: '',
    cardType: '',
    branch: '',
    cardBranding: '',
  };

  const [applicantFormValues, setApplicantFormValues] = React.useState(formValues);
  const [addressFormValues, setAddressFormValues] = React.useState(formValues);
  const [paymentFormValues, setPaymentFormValues] = React.useState(formValues);
  const [selectedDate, setSelectedDate] = React.useState(new Date());
  const [creditCardType, setCreditCardType] = React.useState({ 'cardType': '' });
  const [employment, setEmploymentType] = React.useState({ 'employmentType': '' });

  formValues = {
    firstName: applicantFormValues['firstName'],
    lastName: applicantFormValues['lastName'],
    dateOfBirth: selectedDate,
    phoneNumber: applicantFormValues['phoneNumber'],
    socialSecurityNumber: applicantFormValues['socialSecurityNumber'],
    employmentType: employment['employmentType'],
    occupation: applicantFormValues['occupation'],
    monthlyIncome: applicantFormValues['monthlyIncome'],
    yearsEmployed: applicantFormValues['yearsEmployed'],
    streetAddress: addressFormValues['streetAddress'],
    yearsAtAddress: addressFormValues['yearsAtAddress'],
    city: addressFormValues['city'],
    state: addressFormValues['state'],
    zip: addressFormValues['zip'],
    country: addressFormValues['country'],
    ownership: addressFormValues['ownership'],
    monthlyPayment: addressFormValues['monthlyPayment'],
    cardName: paymentFormValues['cardName'],
    cardType: creditCardType['cardType'],
    branch: paymentFormValues['branch'],
    cardBranding: paymentFormValues['cardBranding'],
  }

  const handleNext = () => {
    if (activeStep === steps.length - 1) {
      console.log("activeStep ===>", activeStep);
      console.log("formValues ===>", formValues);

      axios.post("http://localhost:8090/api/v1/card", {

        firstName: formValues['firstName'],
        lastName: formValues['lastName'],
        dateOfBirth: formValues['dateOfBirth'],
        phoneNumber: formValues['phoneNumber'],
        socialSecurityNumber: formValues['socialSecurityNumber'],
        employmentType: formValues['employmentType'],
        occupation: formValues['occupation'],
        monthlyIncome: formValues['monthlyIncome'],
        yearsEmployed: formValues['yearsEmployed'],
        streetAddress: formValues['streetAddress'],
        yearsAtAddress: formValues['yearsAtAddress'],
        city: formValues['city'],
        state: formValues['state'],
        zip: formValues['zip'],
        country: formValues['country'],
        ownership: formValues['ownership'],
        monthlyPayment: formValues['monthlyPayment'],
        cardName: formValues['cardName'],
        cardType: formValues['cardType'],
        branch: formValues['branch'],
        cardBranding: formValues['cardBranding'],
      }).then(function (response) {
        console.log(response);
      });
    }
    setActiveStep(activeStep + 1);
  };

  const handleBack = () => {
    setActiveStep(activeStep - 1);
  };

  const changeApplicantFormValues = (key, value) => {
    let values = { ...applicantFormValues };
    values[key] = value;
    setApplicantFormValues(values);
  };

  const changeAddressFormValue = (key, value) => {
    let values = { ...addressFormValues };
    values[key] = value;
    setAddressFormValues(values);
  };

  const changePaymentFormValue = (key, value) => {
    let values = { ...paymentFormValues };
    values[key] = value;
    setPaymentFormValues(values);
  };


  return (
    <React.Fragment>
      <CssBaseline />
      <AppBar position="absolute" color="default" className={classes.appBar}>
        <Toolbar>
          <Typography variant="h6" color="inherit" noWrap>
            Community Bank
          </Typography>
        </Toolbar>
      </AppBar>
      <main className={classes.layout}>
        <Paper className={classes.paper}>
          <Typography component="h1" variant="h4" align="center">
            Credit Card Application
          </Typography>
          <Stepper activeStep={activeStep} className={classes.stepper}>
            {steps.map((label) => (
              <Step key={label}>
                <StepLabel>{label}</StepLabel>
              </Step>
            ))}
          </Stepper>
          <React.Fragment>
            {activeStep === steps.length ? (
              <React.Fragment>
                <Typography variant="h5" gutterBottom>
                  Thank you for confirming your information.
                </Typography>
                <Typography variant="subtitle1">
                  Your queue number is #300112. We have emailed your queue confirmation, and will
                  send you an update once your account has been created.
                </Typography>
              </React.Fragment>
            ) : (
              <React.Fragment>
                {
                  activeStep === 0 ? getStepContent(activeStep, applicantFormValues, changeApplicantFormValues, selectedDate, setSelectedDate, employment, setEmploymentType)
                    : activeStep === 1 ? getStepContent(activeStep, addressFormValues, changeAddressFormValue)
                      : activeStep === 2 ? getStepContent(activeStep, paymentFormValues, changePaymentFormValue, null, null, creditCardType, setCreditCardType)
                        : getStepContent(activeStep, formValues)
                }

                {<div className={classes.buttons}>
                  {activeStep !== 0 && (
                    <Button variant="contained" style={{ outline: 'none' }}
                      className={classes.button}
                      onClick={handleBack}
                    >
                      Back
                    </Button>
                  )}
                  <Button style={{ outline: 'none' }}
                    variant="contained"
                    color="primary"
                    onClick={handleNext}
                    className={classes.button}
                  >
                    {activeStep === steps.length - 1 ? 'Confirm' : 'Next'}
                  </Button>
                </div>}
              </React.Fragment>
            )}
          </React.Fragment>
        </Paper>
        <Copyright />
      </main>
    </React.Fragment>
  );
}
