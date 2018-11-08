#include <stdio.h>
#include <math.h>
#include <time.h>
#include <stdlib.h>
#include <pthread.h>
#include <thread>
#include <wiringPi.h>
#include "3DV2.h"  
void Water_INT();


int main(int argc, char const *argv[])
{	
	std::thread Int(Water_INT);
	char input[64];
	char menu;
	int i=0;
	float value;
	wiringPiSetup();
	//wiringPiISR (WaterPin,INT_EDGE_BOTH, &Water_INT) ;
	pullUpDnControl(WaterPin,PUD_DOWN);
	pinMode(8,OUTPUT);
	pinMode(9,OUTPUT);
	
	while(1){
		printf("Menu->\n\tX... \tY... \tAuto y \t Reset xy \t Output\n");
		gets(input);
		sscanf(input,"%c %f",&menu,&value);
		switch(menu){
			case 'X':
			case 'x':
				printf("X Start\n");
				X.__S_PGo__(&value);
				P.c[0]=value;
				break;
			case 'Y':
			case 'y':
				printf("Y Start\n");
				Y.__S_PGo__(&value);
				P.c[1]=value;
				break;
			case 'A':
			case 'a':
				if (value>0)
					{
					Y.__VGet__(value,true,1);
					}else{
					Y.__VGet__(value,false,1);	
					}
				break;
			case 'R':
			case 'r':
				printf("X Y reset to ziro %f\n",value );
				P.c[0]=0;
				P.c[1]=0;
				break;
			case 'O':
			case 'o':
				printf("Out Put generation activated.\n");
				Is_Out_Put=true;
				break;

			default:
			printf("Cant recognize your order\n");
		}



	}
	return 0;
}

void Water_INT(){
	printf("filedone");
	digitalWrite(9,0);
	while(i++<2000){
		digitalWrite(8,1);
		delay(1);
		digitalWrite(8,0);
		delay(1);
		}
	digitalWrite(9,1);
	i=0;
	while(i++<2000){
		digitalWrite(8,1);
		delay(1);
		digitalWrite(8,0);
		delay(1);
		}
	

	while(1){
	delay(100);
	int i=digitalRead(WaterPin);
	if (i==1){
	Is_In_Water=true;
	digitalWrite(7,1);
	}else{
	Is_In_Water=false;
	digitalWrite(7,0);	
	}
	
	}
}






















//////#include <wiringPi.h>
/////////////////////////////////////////////////////////////////////////////////////////////////||
/*

	#define OUTPUT 1
	unsigned int milli=0;
	unsigned int millis(){return milli;}
#ifndef OUTPUT
#define OUTPUT 1
#endif

	void digitalWrite(int pinnum,int val){
		
	}
	int digitalRead(int pin){

	}
	void wiringPiSetup(void){

	}
	void pinMode(int pinnum,int out_put){

	}
	void delayMicroseconds(unsigned int how_long){milli+=1000;}
	void delay(unsigned int how_long){milli++;}


	void MySetUp(void){
		wiringPiSetup();
	
		}
//*/








//_______________________                                                           C L A S S    F U N C T I O N S

//_____________________________________________________________________              P O I N T 

		
		//~Point();
		Point Point::__add__(Point * inp){
			Point holder(c[0]+inp->c[0],c[1]+inp->c[1],c[2]+inp->c[2]);
			return holder;
			}
		Point Point::__Sadd__(Point * p){
			c[0]+=p->c[0];
			c[1]+=p->c[1];
			c[2]+=p->c[2];			
			}
		void Point::__Move__(Point * p){
			c[0]=p->c[0];
			c[1]=p->c[1];
			c[2]=p->c[2];			
			}
		void Point::__Move__(float x1,float x2, float x3){
			c[0]=x1;
			c[1]=x2;
			c[2]=x3;			
			}
		bool Point::__print__(void){
			printf("Point(%f,%f,%f)",c[0],c[1],c[2]);
			return true;
			}
		bool Point::__sdot__(float i){
			c[0] *= i;
			c[1] *= i;
			c[2] *= i;
			return true;
			}
		float Point::__distances__(Point * p){
			float r=pow(pow(c[0]-p->c[0],2)+pow(c[1]-p->c[1],2)+pow(c[2]-p->c[2],2),0.5);
			return r;
			}



//________________________________________________________                   O R I G I N

		bool Origin::FRead(void){
			FILE *myfile;
			myfile= fopen("OriginData","r");
			char buffer[400];
			float h1,h2,h3;
			fgets(buffer,400,myfile);
			if(sscanf(buffer,"C0:%f\tC1:%f\tC2:%ft\t%*s",&h1,&h2,&h3) !=0 ){
					c[0]=h1;							
					c[1]=h2;
					c[2]=h3;
					fclose(myfile);
					return 1;
				}else{
					printf("unformated line @ OriginData file \n");
					fclose(myfile);
					return 0;
				}


			}
		bool Origin::FWrite(void){
			FILE *myfile;
			myfile= fopen("OriginData","w");
			fprintf(myfile, "C0:%f\tC1:%f\tC2:%f",c[0],c[1],c[2]);
			fclose(myfile);
			return true;
			}
		
		bool Origin::Set_Point_As_Origin(Point p){
			c[0]=p.c[0];
			c[1]=p.c[1];
			c[2]=p.c[2];
			return true;
			}


//_________________________________________________________               V E C T O R

		//float a[3];
		
		//~Vector();
		Vector Vector::__FindDeg__(Vector *v){
			Vector vholder(a[0],a[1],a[2]);
			vholder.__MakeUinc__();
			float radres=(float)acos(vholder.__dot__(v)/v->__Magnitude__());
			Vector vholder2=vholder.__Cross__(v);
			vholder2.__MakeUinc__();
			vholder2.__dot__(radres);
			return vholder2;

			}
		bool Vector::__add__(Vector *v){
			a[0]+=v->a[0];
			a[1]+=v->a[1];
			a[2]+=v->a[2];
			return true;
			}
		bool Vector::__MakeUinc__(void){
			float r=pow((pow(a[0],2)+pow(a[1],2)+pow(a[2],2)),0.5);
			a[0]=a[0]/r;
			a[1]=a[1]/r;
			a[2]=a[2]/r;
			return true;
			}
		float Vector::__dot__(Vector * v){
			float res=a[0] * v->a[0]+a[1] * v->a[1]+a[2] * v->a[2];
			return res;
			}
		Vector Vector::__Cross__(Vector * v){
			float h1,h2,h3;
			h1=a[1]*v->a[2]-a[2]*v->a[1];
			h2=a[2]*v->a[0]-a[0]*v->a[2];
			h3=a[0]*v->a[1]-a[1]*v->a[0];
			Vector h(h1,h2,h3);
			return h;
			}
		float Vector::__Magnitude__(void){
			float r=pow((pow(a[0],2)+pow(a[1],2)+pow(a[2],2)),0.5);
			return r;
			}
		Vector Vector::operator* (const float i){
			Vector h(a[0]*i,a[1]*i,a[2]*i);
			return h;
			}
		Vector Vector::operator+ (const Vector * v){
			Vector h(a[0]+v->a[0],a[1]+v->a[1],a[2]+v->a[2]);
			return h;
			}
		Vector Vector::operator- (const Vector * v){
			Vector h(a[0]-v->a[0],a[1]-v->a[1],a[2]-v->a[2]);
			return h;
			}
		bool Vector::__dot__(float i){
			a[0]*=i;
			a[1]*=i;
			a[2]*=i;
			return true;
			}
		bool Vector::__print__(void){
			printf("Vector :(%f,%f,%f)",a[0],a[1],a[2]);
			return true;
			}
		bool Vector::__actPoint__(Point *inp){
			inp->c[0]+=a[0];
			inp->c[1]+=a[1];
			inp->c[2]+=a[2];
			}


//_____________________________________________________________________                P L A T E 
		
		Point Plate::__translate__(Point *inp){
			Vector vholder1=v;
			Vector vholder2(inp,&p);
			Point pres;
			vholder1.__MakeUinc__();
			vholder1.__dot__(vholder1.__dot__(&vholder2));
			pres=*inp;
			vholder1.__actPoint__(&pres);
			if(__IsOnPlate__(&pres)) return pres;
			pres=*inp;
			vholder1.__dot__(-1);
			vholder1.__actPoint__(&pres);
			if(__IsOnPlate__(&pres)) return pres;
			else printf("An Error Ocurred\n"); return *inp;
			}
		Line Plate::__translate__(Line *ina){
			Point h[2];
			h[1]=__translate__(&ina->SP);
			h[2]=__translate__(&ina->EP);
			Line res(&h[1],&h[2]);
			return res;
			}
		
		//Curve __translate__(Curve *inc);
		bool Plate::__IsOnPlate__(Point *inp){
			
			//Point p3(1.2,3,6.44);
			Vector vh(inp,&p);
			//p.__print__();
			//v.__print__();
			float h=v.__dot__(&vh);
			if(fabs(h)<MINRESOLUTION){
			//printf("%f__M\n",h);
			return true;	
				}else {return false;}
			}
		bool Plate::__IsOnPlate__(Curve *inc){
			if(__IsOnPlate__(&inc->SP) && __IsOnPlate__(&inc->EP) && __IsOnPlate__(&inc->CP)) return true;
			else return false;
			}
		bool Plate::__IsOnPlate__(Axis *ina){
			if(__IsOnPlate__(&ina->p) && ina->v.__dot__(&v)<MINRESOLUTION) return true;
			else return false;
			}
		

			//Curve Plate::__intersection__(Plate *inp);
			//Plate::~Plate();
		
		
		Point Plate::__intersection__(Line *l){
			
			float x[3];
			float d= v.a[0] * p.c[0] + v.a[1] * p.c[1] + v.a[2] * p.c[2];
			//p.__print__();
			v.__print__();
			//printf("D is%f\n",d);
			l->v.__print__();
			if(v.__dot__(&l->v)!=0 && !__IsOnPlate__(l)){
				if(l->v.a[0] != 0){	
					x[0]=(d+(l->p.c[0]/l->v.a[0])*(v.a[1]*l->v.a[1]+v.a[2]*l->v.a[2])-(v.a[1]*l->p.c[1]+v.a[2]*l->p.c[2]))/(v.a[0]+v.a[1]*l->v.a[1]/l->v.a[0]+v.a[2]*l->v.a[2]/l->v.a[0]);
					x[1]=l->v.a[1]*(x[0]-l->p.c[0])/l->v.a[0] + l->p.c[1];
					x[2]=l->v.a[2]*(x[0]-l->p.c[0])/l->v.a[0] + l->p.c[2];
					Point resP(x[0],x[1],x[2]);
				//	resP.__print__();
					return resP;
				}else if(l->v.a[1] != 0){
					x[0]=l->p.c[0];
					x[1]=(d-v.a[0]*x[0]+v.a[2]*l->v.a[2]*l->p.c[1]/l->v.a[1] - v.a[2]*l->p.c[2])/(v.a[1]+v.a[2]*l->v.a[2]/l->v.a[1]);
					x[2]=l->v.a[2]*(x[1]-l->p.c[1])/l->v.a[1]+l->p.c[2];
					Point resP(x[0],x[1],x[2]);
				//	resP.__print__();
					return resP;
				}else if(l->v.a[2] !=0){
					x[0]=l->p.c[0];
					x[1]=l->p.c[1];
					x[2]=(d-v.a[0]*x[0]-v.a[1]*x[1])/v.a[2];
					Point resP(x[0],x[1],x[2]);
				//	resP.__print__();
					return resP;
				}else {
					printf("Error here have defined a line with ziro vector\n");
					return NULL;
				}
			}else{
				printf("The plate and line are parallel\n");
				return NULL;
				}			
			}


		Point Plate::__intersection__(Axis *l){
			
			float x[3];
			float d= v.a[0] * p.c[0] + v.a[1] * p.c[1] + v.a[2] * p.c[2];
			//p.__print__();
			//v.__print__();
			//printf("D is%f\n",d);
			//l->v.__print__();
			if(v.__dot__(&l->v)!=0 && !__IsOnPlate__(l)){
				if(l->v.a[0] != 0){	
					x[0]=(d+(l->p.c[0]/l->v.a[0])*(v.a[1]*l->v.a[1]+v.a[2]*l->v.a[2])-(v.a[1]*l->p.c[1]+v.a[2]*l->p.c[2]))/(v.a[0]+v.a[1]*l->v.a[1]/l->v.a[0]+v.a[2]*l->v.a[2]/l->v.a[0]);
					x[1]=l->v.a[1]*(x[0]-l->p.c[0])/l->v.a[0] + l->p.c[1];
					x[2]=l->v.a[2]*(x[0]-l->p.c[0])/l->v.a[0] + l->p.c[2];
					Point resP(x[0],x[1],x[2]);
				//	resP.__print__();
					return resP;
				}else if(l->v.a[1] != 0){
					x[0]=l->p.c[0];
					x[1]=(d-v.a[0]*x[0]+v.a[2]*l->v.a[2]*l->p.c[1]/l->v.a[1] - v.a[2]*l->p.c[2])/(v.a[1]+v.a[2]*l->v.a[2]/l->v.a[1]);
					x[2]=l->v.a[2]*(x[1]-l->p.c[1])/l->v.a[1]+l->p.c[2];
					Point resP(x[0],x[1],x[2]);
				//	resP.__print__();
					return resP;
				}else if(l->v.a[2] !=0){
					x[0]=l->p.c[0];
					x[1]=l->p.c[1];
					x[2]=(d-v.a[0]*x[0]-v.a[1]*x[1])/v.a[2];
					Point resP(x[0],x[1],x[2]);
				//	resP.__print__();
					return resP;
				}else {
					printf("Error here have defined a Axis with ziro vector\n");
					return NULL;
				}
			}else{
				printf("The plate and Axies are parallel\n");
				return NULL;
				}			
			}

		Axis Plate::__intersection__(Plate *inp){
			Vector linevec =v.__Cross__(&inp->v);
			linevec.__MakeUinc__();
			Point pointholder=__translate__(&inp->p);
			Line lineholder(&p,&pointholder);
			pointholder=inp->__intersection__(&lineholder);
			Axis res(&pointholder,&linevec);
			return res;
			}

		void Plate::__print__(void){
			printf("Plate +_____________________________________________________+\nV:");
			v.__print__();printf("\t P:");
			p.__print__();
			printf("\nPlate +_____________________________________________________+\n");
			}

//____________________________________________________________                                   C U R V E 
		
		//virtual bool Curve::__print__() = 0 ;
		Point Curve::__MidPoint__(){
			Point holder=SP.__add__(&EP);
			holder.__sdot__(0.5);
			return holder;
			}
		bool Curve::__reversdirection__(){
			Point holder;
			holder=SP;
			SP=EP;
			EP=holder;
			v.__dot__(-1);
			}

//_____________________________________________________________________            L I N E 
		
		bool Line::__print__(void){
			printf("Line start From  ");
			SP.__print__();
			printf(" To  ");
			EP.__print__();
			printf("   With direction of :  ");
			v.__print__();
			printf("\n");
			return true;
			}
		Point Line::__intersection__(Line *l){
			v.__MakeUinc__();
			float x[3];
			l->v.__MakeUinc__();
			//printf("%i \t , %f , %f\n",(1-fabs(v.__dot__(&l->v))) < MINRES1-fabs(v.__dot__(&l->v)))OLUTION,1-fabs(v.__dot__(&l->v)),MINRESOLUTION );
			if ((1-fabs(v.__dot__(&l->v))) > MINRESOLUTION){

				Vector vholder=v.__Cross__(&l->v);
				vholder.__MakeUinc__();
				Plate plholder(&p,&vholder);
				if(plholder.__IsOnPlate__(&l->p)){
						Line * lholder=new Line(&p,&l->p);
						Point  midp = lholder->__MidPoint__();
						plholder.v.__actPoint__(&midp);
						Vector * vholder=new Vector(&p,&midp);
						Vector vholder1=vholder->__Cross__(&v);
						vholder=new Vector(&l->p,&midp);
						Vector vholder2=vholder->__Cross__(&l->v);
						Plate p1(&p,&vholder1);
						Plate p2(&l->p,&vholder2);
						Axis a1=p1.__intersection__(&p2);
						Point resP=plholder.__intersection__(&a1);
						return resP;
					}else{
					printf("The Lines are diverses\n");
						}

				}else{
					printf("The Lines are parallel , no intersection\n");
					return NULL;
					}
			}
		bool Line::__IsOnCurve__(Point *inp){
			Vector vholder(&p,inp);
			vholder.__MakeUinc__();
			if(v.a[0]!=0){
				if(inp->c[1]==(v.a[1]*(inp->c[0]-p.c[0])/v.a[0]+p.c[1]) && (1-fabs(vholder.__dot__(&v)))<MINRESOLUTION) return true;
				else return false;
			}else if(v.a[1]!=0){
				if(inp->c[2]==(v.a[2]*(inp->c[1]-p.c[1])/v.a[1]+p.c[2]) && (1-fabs(vholder.__dot__(&v)))<MINRESOLUTION) return true;
				else return false;
			}else{
				if(inp->c[0]==p.c[0] && inp->c[1]==p.c[1] && (1-fabs(vholder.__dot__(&v)))<MINRESOLUTION) return true;
				else return false;
				}

			}
		bool Line::__Extend__(Line *l){
			Point p=__intersection__(l);
			if(SP.__distances__(&p)<EP.__distances__(&p)) {SP=p;}
			else {
				EP=p;
				}
				return true;
			}
		Line * Line::__OffSet__(float value,Plate * Pset){
			Vector vholder=v.__Cross__(&Pset->v);
			vholder.__MakeUinc__();
			vholder.__dot__(value);
			Line * Lres=new Line(&SP,&EP);
			vholder.__actPoint__(&Lres->SP);
			vholder.__actPoint__(&Lres->EP);
			Line * Lr=new Line(&Lres->SP,&Lres->EP);
			return Lr;
			}
		float Line::__length__(void){
			Vector vholder(&SP,&EP);
			return vholder.__Magnitude__();
			}


//_____________________________________________________________________            A R C

		Arc * Arc::__OffSet__(float value){
			//Plate * Pset = new Plate(SP,CP,EP);
			Vector vholder1(&CP,&SP);
			Vector vholder2(&CP,&EP);
			vholder1.__MakeUinc__();
			vholder2.__MakeUinc__();
			vholder1.__dot__(value);
			vholder2.__dot__(value);
			Arc * res = new Arc(&SP,&EP,&CP) ;
			return res;

		}

		bool Arc::__print__(){
			printf("Arc +_____________________________________________________________________+\n");
			printf("With Center of\n");
			CP.__print__();
			printf("And Radus of %f\n Start From \t",__radus__());
			SP.__print__();
			printf(" --to-->   ");
			EP.__print__();
			printf("\nArc +_____________________________________________________________________+\n");
			}
		float Arc::__radus__(){
			Vector vholder(&CP,&SP);
			return vholder.__Magnitude__();
			 }
		Point Arc::__intersection__(Line *l){}
		bool Arc::__IsOnCurve__(Point *inp){}
		float Arc::__length__(void){}
		float Arc::__Rad_arc__(void){
			Vector res;
			Vector vholder1(&CP,&SP);
			Vector vholder2(&CP,&EP);
			vholder1.__MakeUinc__();
			vholder2.__MakeUinc__();
			res=vholder1.__FindDeg__(&vholder2);
			return res.__Magnitude__();

		}

		
//___________________________________________________________________              A X I S 


//___________________________________________________________________ 				Motor
		bool Motor::__S_SGo__(){
			static bool state=true;
			if(state){
				digitalWrite(PulsePin,1);state=false;
				if(direction == 1){
					*Position+=Screw_Pitch/Step_Per_Rev;
					}else if(direction == -1){
					*Position-=Screw_Pitch/Step_Per_Rev;
					}else {
					if(digitalRead(DirPin)==1){
						*Position+=Screw_Pitch/Step_Per_Rev;
						}else{
						*Position-=Screw_Pitch/Step_Per_Rev;
						}
					}
				}else{
				digitalWrite(PulsePin,0);state=true;}
			}

		void __S_PGo__th(Motor * a){
			int i=0;
			//printf("___________________________________________________________________%f,%f\n",fabs(*a->Position - a->DisPosition),(1.5*a->Screw_Pitch/a->Step_Per_Rev));

			while(fabs(*a->Position - a->DisPosition)>(1.5*a->Screw_Pitch/a->Step_Per_Rev)){
				a->__S_SGo__();
				delayMicroseconds(a->MinDelay);	
					if (i++>500)
					{
						printf("\t \t \t \t \t \t -->M1 is in %f <--\n",*a->Position);
						i=0;
					}
				}

			a->S_PGo__=false;
			pthread_exit(NULL);	
			}

		void __S_PGo__th2(Motor * a){
			int i=0;
			unsigned int Mydelay=(unsigned int)1000000/(a->Velocity*a->Step_Per_Rev);
			//printf("___________________________________________________________________%f,%f\n",fabs(*a->Position - a->DisPosition),(1.5*a->Screw_Pitch/a->Step_Per_Rev));

			while(fabs(*a->Position - a->DisPosition)>(1.5*a->Screw_Pitch/a->Step_Per_Rev)){
				a->__S_SGo__();
				delayMicroseconds(Mydelay);	
					if (i++>500)
					{
						printf("\t \t \t \t \t \t -->M1 is in %f <--\n",*a->Position);
						i=0;
					}
				}

			a->S_PGo__=false;
			pthread_exit(NULL);	
			}


		void __S_PGo__th3(Motor * a){
			int i=0;
			unsigned int Mydelay=200;
			unsigned int MyMilli=millis();
			float xholder=(float)(pow(a->Velocity,2)/(2*a->DisAcceleration));

			//printf("___________________________________________________________________%f,%f\n",fabs(*a->Position - a->DisPosition),(1.5*a->Screw_Pitch/a->Step_Per_Rev));

			while(fabs(*a->Position - a->DisPosition) > xholder+(1.5*a->Screw_Pitch/a->Step_Per_Rev)){
				a->__S_SGo__();
				delayMicroseconds(Mydelay);	
					if (i++>100)
					{	
						a->Velocity=(float)(millis() - MyMilli)*a->Acceleration/1000;
						Mydelay=(unsigned int)1000000/(a->Velocity*a->Step_Per_Rev);
						printf("\t \t \t \t \t \t -->M1 is in %f <--\n",*a->Position);
						i=0;
					}
				}


			while(fabs(*a->Position - a->DisPosition) > (1.5*a->Screw_Pitch/a->Step_Per_Rev)){
				a->__S_SGo__();
				delayMicroseconds(Mydelay);	
					if (i++>100)
					{	
						a->Velocity=(float)(millis() - MyMilli)*a->DisAcceleration/1000;
						Mydelay=(unsigned int)1000000/(a->Velocity*a->Step_Per_Rev);
						xholder=(float)(pow(a->Velocity,2)/(2*a->DisAcceleration));
						printf("\t \t \t \t \t \t -->M1 is in %f <--\n",*a->Position);
						i=0;
					}
				}

			a->S_PGo__=false;
			pthread_exit(NULL);	
			}


		void __S_PGo__th4(Motor * a){
			int i=0;
			unsigned int Mydelay=1000000;
			unsigned int MyMilli=millis();
			float xholder=(float)(pow(a->Velocity,2)/(2*a->DisAcceleration));

			//printf("___________________________________________________________________%f,%f\n",fabs(*a->Position - a->DisPosition),(1.5*a->Screw_Pitch/a->Step_Per_Rev));

			while(fabs(*a->Position - a->DisPosition) > xholder+(1.5*a->Screw_Pitch/a->Step_Per_Rev)){
				a->__S_SGo__();
				delayMicroseconds(Mydelay);	
					if (i++>500)
					{	
						if(fabs(a->Velocity) < fabs(a->DisVelocity)){
							a->Velocity=(float)(millis() - MyMilli)*a->Acceleration/1000;
							Mydelay=(unsigned int)1000000/(a->Velocity*a->Step_Per_Rev);
							}
					
						printf("\t \t \t \t \t \t -->M1 is in %f <--\n",*a->Position);
						i=0;
					}
				}


			while(fabs(*a->Position - a->DisPosition) > (1.5*a->Screw_Pitch/a->Step_Per_Rev)){
				a->__S_SGo__();
				delayMicroseconds(Mydelay);	
					if (i++>500)
					{	
						a->Velocity=(float)(millis() - MyMilli)*a->DisAcceleration/1000;
						Mydelay=(unsigned int)1000000/(a->Velocity*a->Step_Per_Rev);
						printf("\t \t \t \t \t \t -->M1 is in %f <--\n",*a->Position);
						i=0;
					}
				}

			a->S_PGo__=false;
			pthread_exit(NULL);	
			}

		void __S_PGo__th5(Motor * a){
			int i=0;
			unsigned int Mydelay=(unsigned int)1000000/(Y.Velocity*Y.Step_Per_Rev);
			
			while(Y.VGet__){
				if (!Y.VGet__)	break;

				if(digitalRead(WaterPin)==1){
					Y.direction=-1;
					digitalWrite(Y.DirPin,0);
					}else{
					Y.direction=1;
					digitalWrite(Y.DirPin,1);
					}


				
				Y.__S_SGo__();
				delayMicroseconds(Mydelay);	
					if (i++>500)
					{
						printf("\t \t \t \t \t -->Y is in %f <--  %i \n",*Y.Position,Y.VGet__);
						i=0;
					}
				}
			pthread_exit(NULL);	
			}


		 bool Motor::__S_PGo__(float * inPosition){

			if(*Position - *inPosition < 0){direction =1;digitalWrite(DirPin,1);}
			if(*Position - *inPosition > 0){direction =-1;digitalWrite(DirPin,0);}
			DisPosition=*inPosition;
			std::thread mythread(__S_PGo__th,this);
			mythread.join();
			S_PGo__=true;
			}

		bool Motor::__S_PGo__(float * inPosition,float inVelocity){
			if(*Position - *inPosition < 0){direction =1;digitalWrite(DirPin,1);}
			if(*Position - *inPosition > 0){direction =-1;digitalWrite(DirPin,0);}
			DisPosition=*inPosition;
			Velocity=inVelocity;
			std::thread mythread(__S_PGo__th2,this);
			mythread.join();
			S_PGo__=true;
			}
		bool Motor::__S_PGo__(float *inPosition,float instartAcceleration,float inendAcceleration){
			if(*Position - *inPosition < 0){direction =1;digitalWrite(DirPin,1);}
			if(*Position - *inPosition > 0){direction =-1;digitalWrite(DirPin,0);}
			DisPosition=*inPosition;
			Acceleration=instartAcceleration;
			DisAcceleration=DisAcceleration;
			std::thread mythread(__S_PGo__th3,this);
			mythread.join();
			S_PGo__=true;

		}
		bool Motor::__S_PGo__(float *inPosition,float instartAcceleration,float inendAcceleration,float VMax){
			if(*Position - *inPosition < 0){direction =1;digitalWrite(DirPin,1);}
			if(*Position - *inPosition > 0){direction =-1;digitalWrite(DirPin,0);}
			DisPosition=*inPosition;
			Acceleration=instartAcceleration;
			DisAcceleration=DisAcceleration;
			DisVelocity=VMax;
			std::thread mythread(__S_PGo__th4,this);
			mythread.join();
			S_PGo__=true;
		}

		
		bool Motor::__VGet__(float inVelocity,bool StartStop,char indirection){
			VGet__=StartStop;
			printf("%i\n",VGet__);
			if(StartStop && VGet__){
				printf("P2\n");
			
			Velocity=inVelocity;
			direction=indirection;
				if(direction==1)
					digitalWrite(DirPin,1);
				if (direction==-1)
					digitalWrite(DirPin,0);	
			std::thread mythread(__S_PGo__th5,this);
			mythread.join();
			
			}else if (StartStop && direction != indirection){
				printf("P3\n");
				direction=indirection;
				if(direction==1)
					digitalWrite(DirPin,1);
				if (direction==-1)
					digitalWrite(DirPin,0);	
			}else{
			printf("P4\n");
			VGet__=false;	
				}
		}

		bool Motor::__S_PSmoothGo__(float *inPosition, float SmoothCoef){}
		bool Motor::__AGet__(float inAcceleration){}

		void __output__th(Motor * a){
			FILE * mylog;
			/*mylog=
			
			
			while(1){
				a->__S_SGo__();
				delayMicroseconds(Mydelay);	
					if (i++>100)
					{
						printf("\t \t \t \t \t \t -->M1 is in %f <--\n",*a->Position);
						i=0;
					}
				}
			pthread_exit(NULL);	*/
			}

		bool Motor::__output__(){

			}



		bool Motor::__Proc__(){}
/*
		float Screw_Pitch;
		unsigned int Step_Per_Rev;
		int PulsePin,DirPin;
		float Position;
		float Velocity;
		float Acceleration;
		char direction;
		bool __S_SGo__;
		bool __S_PGo__;
		bool __S_PSmoothGo__;
		bool __VGet__;
		bool __AGet__;
		bool _is_need_proc_;

		float DisPosition;
		float DisVelocity;
		float DisAcceleration;

		usgined int mindelay=maxvel*Step_per_Rev/5

*/