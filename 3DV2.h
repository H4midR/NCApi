#include <stdio.h>
#include <math.h>
#include <time.h>
#include <stdlib.h>
#include <pthread.h>
#include <thread>  
#include <wiringPi.h>

/*


	void digitalWrite(int pinnum,int val);
	int digitalRead(int pin);
	void wiringPiSetup(void);
	void pinMode(int pinnum,int out_put);
	void delayMicroseconds(unsigned int how_long);
	void delay(unsigned int how_long);


//*/

#define MINRESOLUTION 0.000005
#define R2D 57.29577951
#define D2R 0.017453292
#define WaterPin 8



//													3D Space Classes
class Space;
class Origin;
class Point;
class Vector;
class Curve;
class Line;
class Arc;
class Chain;
class Plate;
class Axis;
class Motor;
class SteperMotor;
class ServoMotor;


class MFile;


class Point
	{
	public:
		float c[3];
		Point(){
			c[0]=0;
			c[1]=0;
			c[2]=0;
			}
		Point(float *inc){
			c[0]=inc[0];
			c[1]=inc[1];
			c[2]=inc[2];
			}
		Point(float x0,float x1,float x2){
			c[0]=x0;
			c[1]=x1;
			c[2]=x2;
			}
		//~Point();
		Point __add__(Point * inp);
		Point __Sadd__(Point * p);
		void __Move__(Point * p);
		void __Move__(float x1,float x2, float x3);
		bool __print__(void);
		bool __sdot__(float i);
		float __distances__(Point * p);
		/*char * GetS(void){
			char f[200];
			sprintf(f,"%f\t%f\t%f",c[0],c[1],c[2]);
			return f;
			}*/
	};

class Origin
	{
	public:
		float c[3];
		bool FRead(void);
		bool FWrite(void);
		Origin(){
			c[0]=0;
			c[1]=0;
			c[2]=0;
			}
		Origin(bool Is_read_from_file){

			}
		~Origin(){}
		bool Set_Point_As_Origin(Point p);
	};

class Vector
	{
	public:
		float a[3];
		Vector(){
			a[0]=0;
			a[1]=0;
			a[2]=0;
			}	
		Vector(Point * p){
			a[0]=p->c[0];
			a[1]=p->c[1];
			a[2]=p->c[2];	
			}
		Vector(Point * p1,Point * p2){
			a[0]=p2->c[0]-p1->c[0];
			a[1]=p2->c[1]-p1->c[1];
			a[2]=p2->c[2]-p1->c[2];	
			}
		Vector(float x1,float x2,float x3){
			a[0]=x1;
			a[1]=x2;
			a[2]=x3;	
			}
		//~Vector();
		Vector __FindDeg__(Vector * v);
		bool __add__(Vector *v);
		bool __MakeUinc__(void);
		float __dot__(Vector * v);
		Vector __Cross__(Vector * v);
		float __Magnitude__(void);
		Vector operator* (const float i);
		Vector operator+ (const Vector * v);
		Vector operator- (const Vector * v);
		bool __dot__(float i);
		bool __print__(void);
		bool __actPoint__(Point *inp);
	};

class Plate
	{
	public:
		Point p;
		Vector v;
		Plate(){}
		Plate(Point *inp , Vector *inv){
			p=*inp;
			v=*inv;
			v.__MakeUinc__();
			}
		Plate(Point * p1,Point * p2, Point * p3){
			Vector hv1(p1,p2);
			Vector hv2(p1,p3);
			v=hv1.__Cross__(&hv2);
			v.__MakeUinc__();
			//v.__print__();
			p=*p1;
			}

		bool __IsOnPlate__(Point *inp);
		bool __IsOnPlate__(Curve *inc);
		bool __IsOnPlate__(Axis *ina);
		Point __translate__(Point *inp);
		Line __translate__(Line *ina);
		Point __intersection__(Line *l);
		Point __intersection__(Axis *l);
		Axis __intersection__(Plate *inp);
		void __print__(void);
		//~Plate();
	};

class Axis
	{
	public:
		Point p;
		Vector v;
		Axis(){}
		Axis(Point *inp,Vector *dir){
			p=*inp;
			v=*dir;
			v.__MakeUinc__();
			}
		Axis(Point p1,Point p2){
			Vector hv(&p1,&p2);
			v=hv;
			v.__MakeUinc__();
			p=p1;
			}
		//~Axis();

	};

class Curve : public Axis
	{
	public:
		Point SP;
		Point EP;
		Point CP;
		Curve():Axis(){}
		Curve(Point *p1,Point *p2):Axis(*p1,*p2){
			SP=*p1;
			EP=*p2;
			CP=*p2;
			}
		Curve(Point *p1,Point *p2,Point *pc):Axis(*p1,*p2){
			SP=*p1;
			EP=*p2;
			CP=*pc;
			}
		/*Curve(Point *p1,Vector *v1):Axis(*p1,*v1){

			}*/
		virtual bool __print__() = 0 ;
		virtual Point __intersection__(Line *l) = 0;
		virtual bool __IsOnCurve__(Point *inp) = 0;
		virtual Line * __OffSet__(float value,Plate * Pset) = 0;
		virtual Arc * __OffSet__(float value) = 0;
		virtual float __radus__(void) = 0;
		virtual float __length__(void) = 0;
		virtual float __Rad_arc__(void) = 0;
		Point __MidPoint__();
		bool __reversdirection__();

		//~Curve();
		
	};

class Line : public Curve
	{
	public:
		Line() : Curve() {}
		Line(Point * inp1,Point * inp2) : Curve(inp1,inp2) {}
		Point __intersection__(Line *l);
		bool __IsOnCurve__(Point *inp);
		bool __print__(void);
		bool __Extend__(Line *l);
	    Line * __OffSet__(float value,Plate * Pset);
	    Arc * __OffSet__(float value){}
	    float __radus__(void){return 1e30;}
	    float __length__(void);
	    float __Rad_arc__(void){return 0;}
		//~Line();
		
	};

class Arc : public Curve
	{
	public:
		Arc():Curve(){}
		Arc(Point *ps,Point *pe,Point *pc):Curve(ps,pe,pc){}
		Arc(Point *ps,Point *pe, float R,Plate *inplate):Curve(ps,pe){
			Vector vholder=v;
			vholder.__Cross__(& inplate->v);
			vholder.__MakeUinc__();
			Line lholder(ps,pe);
			CP=lholder.__MidPoint__();
			vholder.__dot__(R);
			vholder.__actPoint__(&CP);
			}
		Arc * __OffSet__(float value);
		Line * __OffSet__(float value,Plate * Pset){}
		bool __print__();
		Point __intersection__(Line *l);
		bool __IsOnCurve__(Point *inp);
		float __radus__();
		float __length__(void);
		float __Rad_arc__(void);


		//~Arc();
	};

class Motor
	{
	public:
		float Screw_Pitch;
		unsigned int Step_Per_Rev;
		int PulsePin,DirPin;
		float * Position;
		float Velocity;
		float Acceleration;
		char direction;
		bool S_SGo__;
		bool S_PGo__;
		bool S_PSmoothGo__;
		bool VGet__;
		bool AGet__;
		bool _is_need_proc_;
		unsigned int MinDelay;

		float DisPosition;
		float DisVelocity;
		float DisAcceleration;
		std::thread mythreads;
		FILE * LogFile;

		Motor(char * filename,float * PPos){
			FILE * myfile;
			float h;
			myfile= fopen(filename,"r");
			fscanf(myfile,"Step_Per_Rev %i\nScrew_Pitch %f\nPulsePin %i\nDirPin %i\nMax RPSec %f",&Step_Per_Rev,&Screw_Pitch,&PulsePin,&DirPin,&h);
			printf("Motor set as :: %i\t%f\t%i\t%i\t%f\n",Step_Per_Rev,Screw_Pitch,PulsePin,DirPin,h);
			fclose(myfile);
			LogFile=fopen("%sLogFile.txt","w+");
			MinDelay=(unsigned int)1000000/(h*Step_Per_Rev);
			printf("%u\n",MinDelay);
			Position=PPos;
			printf("this Motor is at %f\n",*Position);
			pinMode(DirPin,OUTPUT);
			pinMode(PulsePin,OUTPUT);
			}
		bool __S_SGo__();
		bool __S_PGo__(float * inPosition);
	//	void * __S_PGo__th(void * threadid);
		bool __S_PGo__(float * inPosition,float inVelocity);
		bool __S_PGo__(float *inPosition,float instartAcceleration,float inendAcceleration);
		bool __S_PGo__(float *inPosition,float instartAcceleration,float inendAcceleration,float VMax);
		bool __S_PSmoothGo__(float *inPosition, float SmoothCoef);
		bool __VGet__(float inVelocity,bool StartStop,char indirection);
		bool __AGet__(float inAcceleration);
		bool __output__();

		bool __Proc__();

		//~Motor();
		
	};



class MFile
	{		

	private:
		FILE *myfile;
	public:
		bool is_file_finished;
		MFile(char *FileName){
			myfile= fopen(FileName,"rw");
			}
		~MFile(){
			fclose(myfile);
			}
		bool WriteS(char *buffer);
		/*char *ReadS(void){
				static int myN=0;
				char buffer[500];
				static int ln=0;
				if(is_file_finished==1) return 0;
				fgets(buffer,500,myfile);
				ln++;
				if(feof(myfile)) is_file_finished=1;
				//printf("%s\n",buffer );
				if(fgets(buffer,500,myfile)!=NULL){
						myN++;
						return buffer;
					}else{
						printf("unformated line @ %i\n",ln);
						return 0;
					}	
			}*/

	};


bool Is_In_Water=false;
bool Is_Out_Put=false;

Point P;// as the machine position
Motor X("./M1",&P.c[0]);// as x
Motor Y("./M1",&P.c[1]);// as y
//P.c[2] as rpm


