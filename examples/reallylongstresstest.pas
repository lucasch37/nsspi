program InterpreterStressTest;

var
    i, j, k: integer;
    a, b, c: integer;
    r1, r2, r3: real;
    flag1, flag2, flag3: boolean;
    s1, s2, s3: string;

function MaxInt(x: integer; y: integer): integer;
begin
    if x > y then
        MaxInt := x
    else
        MaxInt := y;
end;

function MinInt(x: integer; y: integer): integer;
begin
    if x < y then
        MinInt := x
    else
        MinInt := y;
end;

function AbsInt(x: integer): integer;
begin
    if x < 0 then
        AbsInt := 0 - x
    else
        AbsInt := x;
end;

function Power(base: integer; exponent: integer): integer;
var
    n: integer;
    result: integer;
begin
    result := 1;
    n := 0;

    while n < exponent do
    begin
        result := result * base;
        n := n + 1;
    end;

    Power := result;
end;

function Factorial(n: integer): integer;
begin
    if n <= 1 then
        Factorial := 1
    else
        Factorial := n * Factorial(n - 1);
end;

function Fib(n: integer): integer;
begin
    if n <= 0 then
        Fib := 0
    else if n = 1 then
        Fib := 1
    else
        Fib := Fib(n - 1) + Fib(n - 2);
end;

function IsEven(n: integer): boolean;
begin
    IsEven := (n mod 2) = 0;
end;

function IsOdd(n: integer): boolean;
begin
    IsOdd := (n mod 2) <> 0;
end;

function InRange(x: integer; low: integer; high: integer): boolean;
begin
    if x >= low then
    begin
        if x <= high then
            InRange := true
        else
            InRange := false;
    end
    else
        InRange := false;
end;

function Choose(flag: boolean; x: integer; y: integer): integer;
begin
    if flag then
        Choose := x
    else
        Choose := y;
end;

procedure PrintIntInfo(n: integer);
begin
    writeln('Number: ', n);
    writeln('  even = ', IsEven(n));
    writeln('  odd  = ', IsOdd(n));
    writeln('  abs  = ', AbsInt(n));
end;

procedure ComplexCalculation(x: integer; y: integer);
var
    sum: integer;
    product: integer;
    difference: integer;
    quotient: integer;
    remainder: integer;
begin
    sum := x + y;
    difference := x - y;
    product := x * y;

    if y <> 0 then
        quotient := x div y
    else
        quotient := 0;

    if y <> 0 then
        remainder := x mod y
    else
        remainder := 0;

    writeln('--- Calculation ---');
    writeln('x = ', x);
    writeln('y = ', y);
    writeln('x + y = ', sum);
    writeln('x - y = ', difference);
    writeln('x * y = ', product);
    writeln('x div y = ', quotient);
    writeln('x mod y = ', remainder);
end;

begin

    writeln('====================================');
    writeln('      PASCAL INTERPRETER TEST');
    writeln('====================================');

    { Basic assignments }

    a := 15;
    b := 4;
    c := 0 - 7;

    r1 := 3.14159;
    r2 := 2.5;
    r3 := r1 * r2 + 10.0 / 2.0;

    s1 := 'Hello';
    s2 := 'Pascal';
    s3 := s1 + ' ' + s2 + '!';

    flag1 := true;
    flag2 := false;

    writeln('');
    writeln('--- Basic Types ---');
    writeln('a = ', a);
    writeln('b = ', b);
    writeln('c = ', c);

    writeln('r1 = ', r1);
    writeln('r2 = ', r2);
    writeln('r3 = ', r3);

    writeln('s1 = ', s1);
    writeln('s2 = ', s2);
    writeln('s3 = ', s3);

    writeln('flag1 = ', flag1);
    writeln('flag2 = ', flag2);

    { Arithmetic precedence }

    writeln('');
    writeln('--- Operator Precedence ---');

    a := 2 + 3 * 4;
    writeln('2 + 3 * 4 = ', a);

    a := (2 + 3) * 4;
    writeln('(2 + 3) * 4 = ', a);

    a := 100 - 5 * 6 + 3;
    writeln('100 - 5 * 6 + 3 = ', a);

    a := 100 div 5 + 2 * 3 - 4;
    writeln('100 div 5 + 2 * 3 - 4 = ', a);

    a := 17 mod 5;
    writeln('17 mod 5 = ', a);

    a := 0 - 10 + 3 * 4;
    writeln('0 - 10 + 3 * 4 = ', a);

    { Comparison operators }

    writeln('');
    writeln('--- Comparisons ---');

    writeln('10 = 10: ', 10 = 10);
    writeln('10 <> 5: ', 10 <> 5);
    writeln('10 < 20: ', 10 < 20);
    writeln('20 > 10: ', 20 > 10);
    writeln('10 <= 10: ', 10 <= 10);
    writeln('10 >= 10: ', 10 >= 10);

    { Boolean expressions without AND/OR }

    writeln('');
    writeln('--- Boolean Tests ---');

    flag1 := true;
    flag2 := false;

    flag3 := true;
    writeln('assigned true: ', flag3);
    flag3 := false;
    writeln('assigned false: ', flag3);

    if flag1 then
        flag3 := false
    else
        flag3 := true;
    writeln('inverse of true = ', flag3);

    if flag2 then
        flag3 := false
    else
        flag3 := true;
    writeln('inverse of false = ', flag3);

    if (5 > 3) then
    begin
        if (10 <> 20) then
        begin
            if 1 <= 2 then
                flag3 := true
            else
                flag3 := false;
        end
        else
            flag3 := false;
    end
    else
        flag3 := false;

    writeln('nested boolean test = ', flag3);

    { String operations }

    writeln('');
    writeln('--- Strings ---');

    s1 := 'Hello';
    s2 := 'World';

    s3 := s1 + ', ' + s2 + '!';
    writeln(s3);

    writeln('String equality: ', s1 = 'Hello');
    writeln('String inequality: ', s1 <> s2);

    s1 := 'abc';
    s2 := 'abc';

    if s1 = s2 then
        writeln('Equal strings detected')
    else
        writeln('String comparison failed');

    { Functions }

    writeln('');
    writeln('--- Functions ---');

    writeln('Max(10, 20) = ', MaxInt(10, 20));
    writeln('Min(10, 20) = ', MinInt(10, 20));
    writeln('Abs(0 - 123) = ', AbsInt(0 - 123));
    writeln('Power(2, 10) = ', Power(2, 10));
    writeln('Power(3, 5) = ', Power(3, 5));
    writeln('Factorial(5) = ', Factorial(5));
    writeln('Factorial(8) = ', Factorial(8));

    { Recursive Fibonacci }

    writeln('');
    writeln('--- Fibonacci ---');

    i := 0;

    while i <= 12 do
    begin
        writeln('Fib(', i, ') = ', Fib(i));
        i := i + 1;
    end;

    { Procedures }

    writeln('');
    writeln('--- Procedures ---');

    PrintIntInfo(42);
    PrintIntInfo(0 - 17);

    ComplexCalculation(100, 7);

    { Swap performed inline because reference parameters are not implemented }

    writeln('');
    writeln('--- Swap ---');

    a := 123;
    b := 456;

    writeln('Before swap: a=', a, ' b=', b);

    c := a;
    a := b;
    b := c;

    writeln('After swap:  a=', a, ' b=', b);

    { Nested loops }

    writeln('');
    writeln('--- Nested For Loops ---');

    for i := 1 to 5 do
    begin
        for j := 1 to 5 do
        begin
            k := i * j;
            writeln(i, ' * ', j, ' = ', k);
        end;
    end;

    { More complicated nested loops }

    writeln('');
    writeln('--- Multiplication Pattern ---');

    for i := 1 to 10 do
    begin
        s1 := '';

        for j := 1 to i do
        begin
            s1 := s1 + '*';
        end;

        writeln(s1);
    end;

    { While loop }

    writeln('');
    writeln('--- While Loop ---');

    i := 0;
    a := 0;

    while i < 20 do
    begin
        if a < 100 then
        begin
            a := a + i * 2 - 1;

            if a mod 3 = 0 then
                writeln(
                    'i=', i,
                    ' a=', a,
                    ' divisible by 3'
                )
            else
                writeln(
                    'i=', i,
                    ' a=', a,
                    ' not divisible by 3'
                );

            i := i + 1;
        end
        else
            i := 20;
    end;

    { Conditional function }

    writeln('');
    writeln('--- Conditional Function ---');

    writeln(Choose(true, 100, 200));
    writeln(Choose(false, 100, 200));

    { Range tests }

    writeln('');
    writeln('--- Range Tests ---');

    for i := 0 - 2 to 12 do
    begin
        if InRange(i, 0, 10) then
            writeln(i, ' is inside range')
        else
            writeln(i, ' is outside range');
    end;

    { Heavy expression testing }

    writeln('');
    writeln('--- Heavy Expressions ---');

    a := 5;
    b := 8;
    c := 3;

    r1 := 2.5;
    r2 := 4.0;
    r3 := 10.0;

    if a < b then
    begin
        if b > c then
        begin
            flag1 := true;
        end
        else
            flag1 := false;
    end
    else
        flag1 := false;

    writeln('compound condition = ', flag1);

    a :=
        ((a + b) * c)
        - ((b - c) * a)
        + Power(2, 5);

    writeln('complex integer expression = ', a);

    r3 :=
        (r1 + r2) * 2.0
        - r1 / r2
        + 10.5;

    writeln('complex real expression = ', r3);

    { Deep nesting }

    writeln('');
    writeln('--- Deep Nesting ---');

    i := 1;

    while i <= 5 do
    begin
        j := 1;

        while j <= 5 do
        begin
            if (i + j) mod 2 = 0 then
            begin
                if i > j then
                    writeln('even sum, i > j')
                else if i < j then
                    writeln('even sum, i < j')
                else
                    writeln('even sum, i = j');
            end
            else
            begin
                if i > j then
                    writeln('odd sum, i > j')
                else if i < j then
                    writeln('odd sum, i < j')
                else
                    writeln('odd sum, i = j');
            end;

            j := j + 1;
        end;

        i := i + 1;
    end;

    { Nested function calls }

    writeln('');
    writeln('--- Nested Function Calls ---');

    a :=
        MaxInt(
            MinInt(10, 20),
            AbsInt(0 - 50)
        );

    writeln('nested = ', a);

    a :=
        Power(
            MaxInt(2, 3),
            MinInt(4, 5)
        );

    writeln(
        'Power(Max(2,3), Min(4,5)) = ',
        a
    );

    { Lots of function calls inside expressions }

    writeln('');
    writeln('--- Function Expression Stress ---');

    a :=
        MaxInt(
            Power(2, 3),
            Factorial(4)
        );

    writeln('Max(Power(2,3), Factorial(4)) = ', a);

    b :=
        MinInt(
            AbsInt(0 - 100),
            Power(3, 4)
        );

    writeln('Min(Abs(0 - 100), Power(3,4)) = ', b);

    c :=
        Factorial(
            MinInt(5, 4)
        );

    writeln('Factorial(Min(5,4)) = ', c);

    { Repeated modification }

    writeln('');
    writeln('--- Variable Modification ---');

    a := 0;

    for i := 1 to 10 do
    begin
        a := a + i;

        if IsEven(i) then
            a := a + Power(i, 2)
        else
            a := a - i;

        writeln('i=', i, ' a=', a);
    end;

    { Final stress calculation }

    writeln('');
    writeln('--- Final Stress Calculation ---');

    a := 0;

    for i := 1 to 20 do
    begin
        for j := 1 to 10 do
        begin
            if IsEven(i) then
                a := a + i * j
            else
                a := a - i + j;

            if InRange(a, 0 - 100, 1000) then
                flag1 := true
            else
                flag1 := false;

            if i < j then
                flag2 := true
            else
            begin
                if i = j then
                begin
                    if IsEven(i) then
                        flag2 := true
                    else
                        flag2 := false;
                end
                else
                    flag2 := false;
            end;

            if flag1 then
            begin
                if flag2 then
                    flag3 := true
                else
                    flag3 := false;
            end
            else
                flag3 := false;
        end;
    end;

    writeln('Final a = ', a);
    writeln('Final flag1 = ', flag1);
    writeln('Final flag2 = ', flag2);
    writeln('Final flag3 = ', flag3);

    { Final arithmetic tests }

    writeln('');
    writeln('--- Final Arithmetic Tests ---');

    a := 10 + 20;
    writeln('10 + 20 = ', a);

    a := 50 - 17;
    writeln('50 - 17 = ', a);

    a := 7 * 8;
    writeln('7 * 8 = ', a);

    a := 100 div 4;
    writeln('100 div 4 = ', a);

    a := 100 mod 7;
    writeln('100 mod 7 = ', a);

    a := 0 - (0 - 25);
    writeln('0 - (0 - 25) = ', a);

    r1 := 1.5 + 2.5;
    writeln('1.5 + 2.5 = ', r1);

    r1 := 10.0 - 3.25;
    writeln('10.0 - 3.25 = ', r1);

    r1 := 2.5 * 4.0;
    writeln('2.5 * 4.0 = ', r1);

    r1 := 10.0 / 4.0;
    writeln('10.0 / 4.0 = ', r1);

    { Final strings }

    writeln('');
    writeln('--- Final String Tests ---');

    s1 := 'Pascal';
    s2 := 'Interpreter';
    s3 := s1 + ' ' + s2 + ' Stress Test';

    writeln(s3);

    s1 := '';
    s1 := s1 + 'A';
    s1 := s1 + 'B';
    s1 := s1 + 'C';
    s1 := s1 + 'D';
    s1 := s1 + 'E';

    writeln('Built string: ', s1);

    { Final boolean tests }

    writeln('');
    writeln('--- Final Boolean Tests ---');

    flag1 := true;
    flag2 := false;

    if flag1 then
        writeln('true comparison works')
    else
        writeln('true comparison failed');

    if flag2 then
        writeln('false condition failed')
    else
        writeln('false condition works');

    if flag1 then
        writeln('true condition works')
    else
        writeln('true condition failed');

    if flag2 then
        writeln('false ELSE branch failed')
    else
        writeln('false ELSE branch works');

    writeln('');
    writeln('====================================');
    writeln('       INTERPRETER TEST DONE');
    writeln('====================================');

end.
