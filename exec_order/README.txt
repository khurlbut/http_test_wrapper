This code demonstrates order of execution. I did this explore because I wanted to figure out how to pass flags in as parameters for 
asychronouse tests (time/delay). One of the interesting things this code shows is that Describe and Context execute before
anything other methods. Only variables based on known values at startup may be used for parameters to Describe and Context.

Init() and even BeforeSuite execute AFTER Describe and Context.

This is important to remember when attempting to use FLAGS to control test behavior. The earliest a flag parameter can be evaluated
is in the init() method, meaning you can only reference them from within It() segments.
--- ExecOrderDescribe ---
--- ExecOrderContext---
--- init ---
Running Suite: ExecOrder Suite
==============================
Random Seed: 1541440859
Will run 1 of 1 specs

--- BeforeSuite ---
--- TopLevelFunc ---
--- BeforeEach ---
--- TopLevelFunc ---
--- ExecOrderIt ---
•
Ran 1 of 1 Specs in 0.000 seconds
SUCCESS! -- 1 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 1.477841747s
Test Suite Passed
